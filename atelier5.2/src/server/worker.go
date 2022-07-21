package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"formation-go/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type RegisterCoordinator interface {
	//Register : url is the url of the worker
	Register(url string) error
}

type RegisterCoordinatorFromUrl struct {
	url string
}

func NewRegisterCoordinator(urlCoordinator string) RegisterCoordinator {
	return RegisterCoordinatorFromUrl{url: urlCoordinator}
}

func (rc RegisterCoordinatorFromUrl) Register(url string) error {
	resp, err := http.Post(fmt.Sprintf("%s/register", rc.url), "application/json", strings.NewReader(fmt.Sprintf("{\"url\":\"%s\"}", url)))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("impossible to connect coordinator")
	}
	return nil
}

type Worker struct {
	server     *http.Server
	ackManager model.Ack
	asyncMode  bool
	register   RegisterCoordinator
	port       int
}

func NewWorker(port int, ackManager model.Ack, register RegisterCoordinator, asyncMode bool) Worker {
	server := http.ServeMux{}
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &server,
	}
	w := Worker{&wrapServer, ackManager, asyncMode, register, port}

	server.HandleFunc("/status", w.status)
	server.HandleFunc("/tasks", w.manageTasks)
	server.HandleFunc("/tasks/", w.manageTasks)

	return w
}

func (work Worker) Run() {
	err := work.register.Register(fmt.Sprintf("http://localhost:%d", work.port))
	if err != nil {
		log.Fatal("Impossible to start server", err)
	}
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
	status := model.TaskFinish
	if !task.Do() {
		status = model.TaskError
	}
	work.ackManager.Do(task.Id(), string(status))
}

func (work Worker) status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not allowed method", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("up"))
}
