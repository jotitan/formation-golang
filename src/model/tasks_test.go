package model

import (
	"formation-go/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	message := "Hello Golfi"
	Print{Message: message}.Do()
	assert.True(t, logger.Log.CheckMessage(message))
}

func TestResize(t *testing.T) {
	resize := Resize{
		OriginPath: "/home/img.jpg",
		TargetPath: "/home/img_resize.jpg",
		Height:     600,
		Width:      400,
	}
	resize.Do()
	assert.True(t, logger.Log.CheckMessage("Run resize /home/img.jpg, /home/img_resize.jpg, 600px, 400px"))
}
