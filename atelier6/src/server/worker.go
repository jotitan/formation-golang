package server

import (
	"context"
	"errors"
	"fmt"
	"formation-go/model"
	"github.com/gin-gonic/gin"
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
	server := gin.Default()
	server.HandleMethodNotAllowed = true
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server,
	}
	w := Worker{&wrapServer, ackManager, asyncMode, register, port}

	server.GET("/status", w.status)
	server.POST("/tasks/:id", w.manageTasks)

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

func (work Worker) manageTasks(ctx *gin.Context) {
	idTask, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("unknown id"))
		return
	}
	var payload map[string]interface{}
	ctx.BindJSON(&payload)

	task, err := model.DetectAndCreateTask(payload, func() int { return idTask })
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if work.asyncMode {
		go work.runTask(task)
	} else {
		work.runTask(task)
	}
	ctx.Status(http.StatusOK)
}

func (work Worker) runTask(task model.Task) {
	status := model.TaskFinish
	if !task.Do() {
		status = model.TaskError
	}
	work.ackManager.Do(task.Id(), string(status))
}

func (work Worker) status(ctx *gin.Context) {
	ctx.String(http.StatusOK, "up")
}
