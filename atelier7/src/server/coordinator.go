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
)

type Coordinator struct {
	server      *http.Server
	taskManager *model.Manager
	pool        *PoolWorker
}

func NewCoordinator(port int, taskManager *model.Manager, poolWorker *PoolWorker) Coordinator {
	server := gin.Default()
	server.HandleMethodNotAllowed = true
	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server,
	}
	c := Coordinator{&wrapServer, taskManager, poolWorker}

	server.GET("/status", c.status)
	server.GET("/tasks", c.getTasks)
	server.POST("/tasks", c.addTask)
	server.POST("/tasks/:id", c.updateStatusTask)
	server.GET("/tasks/:id", c.getTask)
	server.POST("/register", c.addWorkerToPool)

	return c
}

func (c Coordinator) Run() {
	log.Println("Start coordinator")
	c.server.ListenAndServe()
}

func (c Coordinator) Stop() {
	c.server.Shutdown(context.Background())
}

func (c Coordinator) status(ctx *gin.Context) {
	ctx.String(200, "up")
}

func (c Coordinator) getTasks(ctx *gin.Context) {
	tasks := c.taskManager.GetAllWithStatus()
	lightTasks := make([]taskDto, len(tasks))
	for i, t := range tasks {
		lightTasks[i] = taskDto{
			t.Task.Type(),
			t.Task.Id(),
			t.Status,
		}
	}
	ctx.JSON(http.StatusOK, lightTasks)
}

func (c Coordinator) getTask(ctx *gin.Context) {
	idTask, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("bad id"))
		return
	}
	task, status := c.taskManager.GetWithStatus(idTask)
	dto := taskDto{
		TypeTask: task.Type(),
		Id:       task.Id(),
		Status:   status,
	}
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("error marshalling"))
		return
	}
	ctx.JSON(http.StatusOK, dto)
}

func (c Coordinator) updateStatusTask(ctx *gin.Context) {
	idTask, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	payload := extractPayload(ctx)
	c.taskManager.UpdateStatus(idTask, model.TaskStatus(payload["status"].(string)))
	c.pool.poolBridge.ReleaseWorker(payload["uuid"].(string))
}

func (c Coordinator) addTask(ctx *gin.Context) {
	payload := extractPayload(ctx)
	task, err := c.taskManager.DetectAndCreateTask(payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.taskManager.Add(task)
	c.pool.Execute(task)
	ctx.String(http.StatusCreated, strconv.Itoa(task.Id()))
}

func (c Coordinator) addWorkerToPool(ctx *gin.Context) {
	payload := extractPayload(ctx)
	c.pool.Register(payload["url"].(string), payload["uuid"].(string), int(payload["capacity"].(float64)))
}

func extractPayload(ctx *gin.Context) map[string]interface{} {
	var payload map[string]interface{}
	ctx.BindJSON(&payload)
	return payload
}

type taskDto struct {
	TypeTask string           `json:"type"`
	Id       int              `json:"id"`
	Status   model.TaskStatus `json:"status"`
}
