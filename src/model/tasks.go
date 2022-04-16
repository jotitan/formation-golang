package model

import (
	"fmt"
	"formation-go/logger"
)

type Task interface {
	Do() bool
	GetId() int
}

type Resize struct {
	Width      int
	Height     int
	OriginPath string
	TargetPath string
	Id         int
}

func (r Resize) Do() bool {
	logger.Log.Println(fmt.Sprintf("Run resize %s, %s, %dpx, %dpx", r.OriginPath, r.TargetPath, r.Height, r.Width))
	return true
}

func (r Resize) GetId() int {
	return r.Id
}

type Print struct {
	Message string
	Id      int
}

func (p Print) Do() bool {
	logger.Log.Println(p.Message)
	return true
}

func (p Print) GetId() int {
	return p.Id
}
