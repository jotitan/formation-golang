package server

import "formation-go/model"

type Coordinator struct{}

func NewCoordinator(port int, manager *model.Manager) Coordinator {
	return Coordinator{}
}

func (c Coordinator) Run() {}

func (c Coordinator) Stop() {}

//lightTask is a dto to give task information
type lightTask struct {
	TypeTask string `json:"type"`
	Id       int    `json:"id"`
	Status   string `json:"status"`
}
