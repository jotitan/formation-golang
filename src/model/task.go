package model

import "fmt"

type Task interface {
	Do() bool
}

type Resize struct {
	Width      int
	Height     int
	OriginPath string
	TargetPath string
}

type Print struct {
	Message string
}

func (p Print) Do() bool {
	fmt.Println(p.Message)
	return true
}
