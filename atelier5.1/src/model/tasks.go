package model

import (
	"fmt"
	"formation-go/logger"
)

type Task interface {
	Do() bool
	Id() int
	Type() string
}

type resize struct {
	width      int
	height     int
	originPath string
	targetPath string
	uuid       int
}

func NewResize(originalPath, targetPath string, height, width, id int) Task {
	return resize{
		originPath: originalPath,
		targetPath: targetPath,
		height:     height,
		width:      width,
		uuid:       id}
}

func (r resize) Type() string {
	return "resize"
}

func (r resize) Do() bool {
	logger.Log.Println(fmt.Sprintf("Run resize %s, %s, %dpx, %dpx", r.originPath, r.targetPath, r.height, r.width))
	return true
}

//Id return unique id of task
func (r resize) Id() int {
	return r.uuid
}

func NewPrint(message string, id int) Task {
	return print{
		message: message,
		uuid:    id,
	}
}

type print struct {
	message string
	uuid    int
}

func (p print) Type() string {
	return "print"
}

func (p print) Do() bool {
	logger.Log.Println(p.message)
	return true
}

//Id return unique id of task
func (p print) Id() int {
	return p.uuid
}
