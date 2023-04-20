package server

import (
	"errors"
	"fmt"
	"formation-go/model"
	"log"
	"net/http"
	"sort"
	"strings"
)

//TaskSenderToWorker define a structure which can send task to a worker
type TaskSenderToWorker interface {
	Send(task model.Task, url string) error
}

type LaunchTask struct{}

func (l LaunchTask) Send(task model.Task, url string) error {
	resp, err := http.Post(fmt.Sprintf("%s/tasks/%d", url, task.Id()), "application/json", strings.NewReader(task.Json()))
	if err == nil && resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("impossible to launch task")
}

type innerWorker struct {
	url      string
	capacity int
	uuid     string
}

type PoolWorker struct {
	workers     map[string]innerWorker
	sender      TaskSenderToWorker
	poolBridge  BridgePoolTask
	chanelTasks chan model.Task
	lastWorker  int
}

func NewWorkerPool(sender TaskSenderToWorker) *PoolWorker {
	return &PoolWorker{
		workers:     make(map[string]innerWorker),
		sender:      sender,
		poolBridge:  BridgePoolTask{make(map[string]*runningWorker)},
		chanelTasks: make(chan model.Task, 10),
		lastWorker:  -1,
	}
}

func (pw *PoolWorker) Size() int {
	return len(pw.workers)
}

func (pw *PoolWorker) Register(url, uuid string, capacity int) bool {
	_, exist := pw.workers[uuid]
	if exist {
		return false
	}
	worker := innerWorker{
		url:      url,
		capacity: capacity,
		uuid:     uuid,
	}
	pw.workers[uuid] = worker
	// Add worker in bridge pool
	pw.poolBridge.AddWorker(worker, pw.chanelTasks)
	log.Println("New worker in pool")
	return true
}

func (pw *PoolWorker) Remove(uuid string) {
	delete(pw.workers, uuid)
	// Remove from pool bridge
}

// Sort by url
func (pw *PoolWorker) getPoolAsList() []innerWorker {
	inners := make([]innerWorker, 0, len(pw.workers))
	for _, inner := range pw.workers {
		inners = append(inners, inner)
	}
	sort.Slice(inners, func(i, j int) bool { return inners[i].url < inners[j].url })
	return inners
}

func (pw *PoolWorker) nextWorker() innerWorker {
	next := (pw.lastWorker + 1) % len(pw.workers)
	pw.lastWorker = next
	inners := pw.getPoolAsList()
	return inners[next]
}

func (pw *PoolWorker) Execute(task model.Task) error {
	if pw.Size() == 0 {
		return errors.New("now worker in the pool")
	}
	pw.chanelTasks <- task
	return nil
	// Old version
	/*inner := pw.nextWorker()
	return pw.sender.Send(task, inner.url)*/
}
