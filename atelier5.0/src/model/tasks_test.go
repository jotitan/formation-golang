package model

import (
	"formation-go/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	// GIVEN
	message := "Hello Golfi"

	// WHEN
	print{message: message}.Do()

	// THEN
	assert.True(t, logger.Log.CheckMessage(message), "Must find message")
}

func TestResize(t *testing.T) {
	// GIVEN
	resize := resize{
		height:     600,
		width:      400,
		originPath: "/home/photo.jpg",
		targetPath: "/home/photo_resize.jpg",
	}

	// WHEN
	resize.Do()

	// THEN
	messageToFind := "Run resize /home/photo.jpg, /home/photo_resize.jpg, 600px, 400px"
	assert.True(t, logger.Log.CheckMessage(messageToFind), "Must find message")
}
