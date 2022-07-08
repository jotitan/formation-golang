package tools

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResize(t *testing.T) {
	// GIVEN
	resizer := LocalResize{}
	folderPath, _ := os.MkdirTemp("", "resize_image")

	path := "../resources/photo_test.jpg"
	currentFolder, _ := os.Getwd()
	if strings.HasSuffix(currentFolder, "tools") {
		path = "../" + path
	}
	data, _ := ioutil.ReadFile(path)

	originalPathFile := filepath.Join(folderPath, "copy_file.jpg")
	resizePathFile := filepath.Join(folderPath, "resize_file.jpg")
	ioutil.WriteFile(originalPathFile, data, os.ModePerm)

	// WHEN
	err := resizer.Resize(originalPathFile, resizePathFile, 400, 600)

	// THEN
	assert.Nil(t, err)

	copyData, _ := ioutil.ReadFile(resizePathFile)
	assert.True(t, len(copyData) < len(data))
}
