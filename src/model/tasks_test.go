package model

import (
	"formation-go/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	message := "Hello my friend"
	Print{message}.Do()
	assert.True(t, logger.Log.CheckMessage(message))
}

func TestResize(t *testing.T) {
	resize := Resize{
		Height:     600,
		Width:      400,
		OriginPath: "/home/photo.jpg",
		TargetPath: "/home/photo_resize.jpg",
	}
	resize.Do()
	assert.True(t, logger.Log.CheckMessage("Run resize /home/photo.jpg, /home/photo_resize.jpg, 600px, 400px"))
}
