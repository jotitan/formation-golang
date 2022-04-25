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
	resizer := LocalResize{}
	folderPath, err := os.MkdirTemp("", "resize_image")
	assert.Nil(t, err)

	path := "../resources/photo_test.jpg"
	currentFolder, _ := os.Getwd()
	if strings.HasSuffix(currentFolder, "tools") {
		path = "../" + path
	}
	data, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	originalPathFile := filepath.Join(folderPath, "copy_file.jpg")
	resizePathFile := filepath.Join(folderPath, "resize_file.jpg")
	ioutil.WriteFile(originalPathFile, data, os.ModePerm)

	err = resizer.Resize(originalPathFile, resizePathFile, 400, 600)
	assert.Nil(t, err)

	copyData, err := ioutil.ReadFile(resizePathFile)
	assert.Nil(t, err)

	assert.True(t, len(copyData) < len(data))
}
