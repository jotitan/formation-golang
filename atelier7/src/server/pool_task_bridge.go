package server

import (
	"errors"
	"fmt"
	"formation-go/model"
)

//var sender TaskSenderToWorker = LaunchTask{}

type requestLimiter chan struct{}

func (rl requestLimiter) add() {
	rl <- struct{}{}
}

func (rl requestLimiter) release() {
	<-rl
}

type runningWorker struct {
	uuid    string
	url     string
	limiter requestLimiter
	sender  TaskSenderToWorker
}

func (running *runningWorker) start(tasks chan model.Task) {
	for {
		running.limiter.add()
		// Read next task only if worker can continue to work
		task := <-tasks
		if running.runTask(task) != nil {
			// If task fail, reinject in channel
			tasks <- task
		}
	}
}

func (running *runningWorker) runTask(task model.Task) error {
	err := running.sender.Send(task, running.url)
	if err != nil {
		// Free limiter
		running.limiter.release()
		return err
	}
	return nil
}

//BridgePoolTask is a bridge which made link between task and the pool of workers to execute
type BridgePoolTask struct {
	runningWorkers map[string]*runningWorker
	sender         TaskSenderToWorker
}

func (bpt *BridgePoolTask) AddWorker(worker innerWorker, chanelTasks chan model.Task) {
	running := &runningWorker{
		uuid:    worker.uuid,
		url:     worker.url,
		sender:  bpt.sender,
		limiter: make(chan struct{}, worker.capacity),
	}
	bpt.runningWorkers[worker.uuid] = running
	go running.start(chanelTasks)
}

// ReleaseWorker release worker (to manage another task)
func (bpt *BridgePoolTask) ReleaseWorker(uuid string) error {
	running, ok := bpt.runningWorkers[uuid]
	if !ok {
		return errors.New(fmt.Sprintf("no worker with id %s", uuid))
	}
	running.limiter.release()
	return nil
}
