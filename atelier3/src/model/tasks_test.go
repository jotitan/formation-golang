package model

import (
	"formation-go/logger"
	"testing"
)

func TestPrint(t *testing.T) {
	message := "Hello Golfi"
	print{message: message}.Do()

	if !logger.Log.CheckMessage(message) {
		t.Error("Must find message", message)
	}
}

func TestResize(t *testing.T) {
	resize := resize{
		height:     600,
		width:      400,
		originPath: "/home/photo.jpg",
		targetPath: "/home/photo_resize.jpg",
	}
	resize.Do()
	messageToFound := "Run resize /home/photo.jpg, /home/photo_resize.jpg, 600px, 400px"
	if !logger.Log.CheckMessage(messageToFound) {
		t.Error("Must find message", messageToFound)
	}
}
