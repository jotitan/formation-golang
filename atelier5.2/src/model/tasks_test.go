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

func TestRealResize(t *testing.T) {
	folder, _ := os.MkdirTemp("", "resize")
	resize := resize{
		height:     600,
		width:      400,
		originPath: filepath.Join("resources", "photo_test.jpg"),
		targetPath: filepath.Join(folder, "output_image.jpeg"),
	}
	if !resize.Do() {
		t.Error("Must resize image")
	}
	testFile, err := os.Open(filepath.Join(folder, "output_image.jpeg"))
	defer testFile.Close()
	if err != nil {
		t.Error("Image is missing")
	}
}
