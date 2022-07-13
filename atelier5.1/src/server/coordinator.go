package server

import (
	"context"
	"encoding/json"
	"fmt"
	"formation-go/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Coordinator struct {
	server      *http.Server
	taskManager *model.Manager
}

func NewCoordinator(port int, taskManager *model.Manager) Coordinator {
	server := http.ServeMux{}
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &server,
	}
	c := Coordinator{&wrapServer, taskManager}

	server.HandleFunc("/status", c.status)
	server.HandleFunc("/tasks", c.manageTasks)
	server.HandleFunc("/tasks/", c.manageTasks)

	return c
}

func (c Coordinator) Run() {
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
	var payload map[string]string
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &payload)
	c.taskManager.UpdateStatus(idTask, payload["status"])
}

func (c Coordinator) addTask(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var payload map[string]interface{}
	err := json.Unmarshal(data, &payload)
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
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(task.Id())))
}

type taskDto struct {
	TypeTask string `json:"type"`
	Id       int    `json:"id"`
	Status   string `json:"status"`
}
