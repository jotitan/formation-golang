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

type Worker struct {
	server     *http.Server
	ackManager model.Ack
	asyncMode  bool
}

func NewWorker(port int, ackManager model.Ack, asyncMode bool) Worker {
	server := http.ServeMux{}
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &server,
	}
	w := Worker{&wrapServer, ackManager, asyncMode}

	server.HandleFunc("/status", w.status)
	server.HandleFunc("/tasks", w.manageTasks)
	server.HandleFunc("/tasks/", w.manageTasks)

	return w
}

func (work Worker) Run() {
	work.server.ListenAndServe()
}

func (work Worker) Stop() {
	work.server.Shutdown(context.Background())
}

func (work Worker) manageTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	idTask, err := strconv.Atoi(strings.ReplaceAll(r.RequestURI, "/tasks/", ""))
	if err != nil {
		http.Error(w, "unknown id", http.StatusBadRequest)
	}
	data, _ := ioutil.ReadAll(r.Body)
	var payload map[string]interface{}
	json.Unmarshal(data, &payload)
	task, err := model.DetectAndCreateTask(payload, func() int { return idTask })
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if work.asyncMode {
		go work.runTask(task)
	} else {
		work.runTask(task)
	}
	w.WriteHeader(http.StatusOK)
}

func (work Worker) runTask(task model.Task) {
	status := "finish"
	if !task.Do() {
		status = "end"
	}
	work.ackManager.Do(task.Id(), status)
}

func (work Worker) status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not allowed method", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("up"))
}
