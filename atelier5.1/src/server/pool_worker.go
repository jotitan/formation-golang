package server

import (
	"formation-go/model"
)

//TaskSenderToWorker define a structure which can send task to a worker
type TaskSenderToWorker interface {
	Send(task model.Task, url string) error
}

type PoolWorker struct {
}

func NewWorkerPool(sender TaskSenderToWorker) *PoolWorker {
	return &PoolWorker{}
}

func (pw *PoolWorker) Add(url string) bool {
	return true
}

func (pw *PoolWorker) Size() int {
	return 0
}

func (pw *PoolWorker) Remove(url string) {
}

func (pw *PoolWorker) Execute(task model.Task) error {
	return nil
}
