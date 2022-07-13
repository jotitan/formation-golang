package server

import (
	"formation-go/model"
)

type Worker struct {
}

func NewWorker(port int, ackManager model.Ack, asyncMode bool) Worker {
	return Worker{}
}

func (work Worker) Run() {
}

func (work Worker) Stop() {
}
