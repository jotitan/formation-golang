package model

import (
	"formation-go/logger"
	"testing"
)

func TestPrint(t *testing.T) {
	message := "Hello my friend"
	Print{message}.Do()

	if !logger.Log.CheckMessage(message) {
		t.Error("Must find message", message)
	}
}

func TestResize(t *testing.T) {
	resize := Resize{
		Height:     600,
		Width:      400,
		OriginPath: "/home/photo.jpg",
		TargetPath: "/home/photo_resize.jpg",
	}
	resize.Do()
	messageToFound := "Run resize /home/photo.jpg, /home/photo_resize.jpg, 600px, 400px"
	if !logger.Log.CheckMessage(messageToFound) {
		t.Error("Must find message", messageToFound)
	}
}
