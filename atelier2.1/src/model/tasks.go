package model

import (
	"fmt"
	"formation-go/logger"
)

type Task interface {
	Do() bool
	Id() int
}

type Resize struct {
	Width      int
	Height     int
	OriginPath string
	TargetPath string
	Uuid       int
}

func (r Resize) Do() bool {
	logger.Log.Println(fmt.Sprintf("Run resize %s, %s, %dpx, %dpx", r.OriginPath, r.TargetPath, r.Height, r.Width))
	return true
}

//Id return unique id of task
func (r Resize) Id() int {
	return r.Uuid
}

type Print struct {
	Message string
	Uuid    int
}

func (p Print) Do() bool {
	logger.Log.Println(p.Message)
	return true
}

//Id return unique id of task
func (p Print) Id() int {
	return p.Uuid
}
