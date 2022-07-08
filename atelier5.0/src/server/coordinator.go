package server

type Coordinator struct{}

func NewCoordinator(port int) Coordinator {
	return Coordinator{}
}

func (c Coordinator) Run() {}

func (c Coordinator) Stop() {}
