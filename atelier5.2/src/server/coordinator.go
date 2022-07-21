package server

import (
	"context"
	"encoding/json"
	"fmt"
	"formation-go/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Coordinator struct {
	server      *http.Server
	taskManager *model.Manager
	pool        *PoolWorker
}

func NewCoordinator(port int, taskManager *model.Manager, poolWorker *PoolWorker) Coordinator {
	server := http.ServeMux{}
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &server,
	}
	c := Coordinator{&wrapServer, taskManager, poolWorker}

	server.HandleFunc("/status", c.status)
	server.HandleFunc("/tasks", c.manageTasks)
	server.HandleFunc("/tasks/", c.manageTasks)
	server.HandleFunc("/register", c.addWorkerToPool)

	return c
}

func (c Coordinator) Run() {
	log.Println("Start coordinator")
	c.server.ListenAndServe()
}

func (c Coordinator) Stop() {
	c.server.Shutdown(context.Background())
}

func (c Coordinator) status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not allowed method", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("up"))
}

func (c Coordinator) manageTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.updateStatusTask(w, r)
	}
	if r.Method == http.MethodGet {
		c.getTask(w, r)
	}
}

func (c Coordinator) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := c.taskManager.GetAllWithStatus()
	lightTasks := make([]taskDto, len(tasks))
	for i, t := range tasks {
		lightTasks[i] = taskDto{
			t.Task.Type(),
			t.Task.Id(),
			t.Status,
		}
	}
	data, _ := json.Marshal(lightTasks)
	w.Write(data)
}

func (c Coordinator) getTask(w http.ResponseWriter, r *http.Request) {
	idTask, err := strconv.Atoi(strings.ReplaceAll(r.RequestURI, "/tasks/", ""))
	if err != nil {
		c.getTasks(w, r)
		return
	}
	task, status := c.taskManager.GetWithStatus(idTask)
	data, err := json.Marshal(taskDto{
		TypeTask: task.Type(),
		Id:       task.Id(),
		Status:   status,
	})
	if err != nil {
		http.Error(w, "error marshalling", http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func (c Coordinator) updateStatusTask(w http.ResponseWriter, r *http.Request) {
	idTask, err := strconv.Atoi(strings.ReplaceAll(r.RequestURI, "/tasks/", ""))
	if err != nil {
		c.addTask(w, r)
		return
	}
	payload, _ := extractPayload(r)
	c.taskManager.UpdateStatus(idTask, model.TaskStatus(payload["status"].(string)))
}

func (c Coordinator) addTask(w http.ResponseWriter, r *http.Request) {
	payload, err := extractPayload(r)
	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
		return
	}
	task, err := c.taskManager.DetectAndCreateTask(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.taskManager.Add(task)
	c.pool.Execute(task)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(task.Id())))
}

func (c Coordinator) addWorkerToPool(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad method", http.StatusMethodNotAllowed)
		return
	}
	payload, _ := extractPayload(r)
	c.pool.Register(payload["url"].(string))
}

func extractPayload(r *http.Request) (map[string]interface{}, error) {
	var payload map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

type taskDto struct {
	TypeTask string           `json:"type"`
	Id       int              `json:"id"`
	Status   model.TaskStatus `json:"status"`
}
