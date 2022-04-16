package model

import (
	"fmt"
	"formation-go/logger"
)

type Task interface {
	Do() bool
}

type Resize struct {
	Width      int
	Height     int
	OriginPath string
	TargetPath string
}

func (r Resize) Do() bool {
	logger.Log.Println(fmt.Sprintf("Run resize %s, %s, %dpx, %dpx", r.OriginPath, r.TargetPath, r.Height, r.Width))
	return true
}

type Print struct {
	Message string
}

func (p Print) Do() bool {
	logger.Log.Println(p.Message)
	return true
}
