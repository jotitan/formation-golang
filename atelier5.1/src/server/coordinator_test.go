package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"formation-go/model"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func startAndWaitServer(manager *model.Manager) (Coordinator, error) {
	server := NewCoordinator(9007, manager)
	go server.Run()
	log.Println("Run server")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		// Call api status
		resp, err := http.Get("http://localhost:9007/status")
		if err == nil && resp.StatusCode == 200 {
			return server, nil
		}
	}
	return Coordinator{}, errors.New("impossible to start server")
}

func TestRunServer(t *testing.T) {
	// GIVEN
	server, err := startAndWaitServer(nil)

	// WHEN - THEN
	assert.Nil(t, err, "Server coordinator must start and be up")
	server.Stop()
}

func TestStatusOnlyGet(t *testing.T) {
	// GIVEN
	server, err := startAndWaitServer(nil)
	assert.Nil(t, err, "Server coordinator must start and be up")

	// WHEN
	resp, _ := http.Post("http://localhost:9007/status", "application/json", bytes.NewBuffer([]byte{}))

	// THEN
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	server.Stop()
}

func TestPostTaskPrint(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	// WHEN
	resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))

	// THEN
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	idTask := string(read(resp.Body))
	assert.Equal(t, fmt.Sprintf("%d", manager.GetAll()[0].Id()), idTask)
	assert.Equal(t, 1, manager.Size())
	server.Stop()
}

func TestPostTaskMultiPrint(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	// WHEN
	_, err = http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))
	_, err = http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"))
	_, err = http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Je sais voler\"}"))

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, 3, manager.Size())
	server.Stop()
}

func TestGetDetailPrintTask(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))
	idTask := string(read(resp.Body))

	// WHEN
	resp, err = http.Get(fmt.Sprintf("http://localhost:9007/tasks/%s", idTask))

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var response map[string]interface{}
	err = json.Unmarshal(read(resp.Body), &response)

	assert.Nil(t, err)
	assert.Equal(t, "print", response["type"])
	assert.Equal(t, idTask, fmt.Sprintf("%.0f", response["id"]))
	assert.Equal(t, "running", response["status"])

	server.Stop()
}

func read(reader io.Reader) []byte {
	data, _ := ioutil.ReadAll(reader)
	return data
}

func TestGetMultiTask(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	manager.Add(model.NewPrint("bonjour mon ami", 1))
	manager.Add(model.NewPrint("je sais voler", 2))
	manager.Add(model.NewPrint("c'est l'anniversaire d'andy", 3))

	// WHEN
	resp, _ := http.Get("http://localhost:9007/tasks")

	// THEN
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var lightTasks []lightTask
	err = json.Unmarshal(read(resp.Body), &lightTasks)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(lightTasks))

	server.Stop()
}

func TestPostTaskResize(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	// WHEN
	payload := "{\"type\":\"resize\",\"path\":\"/file.png\",\"target\":\"/file.png\", \"height\":200, \"width\":300}"
	resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))

	// THEN
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, 1, manager.Size())
	server.Stop()
}

func TestPostTaskFail(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	server, err := startAndWaitServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	// WHEN
	resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"unknown\"}"))

	// THEN
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, 0, manager.Size())
	server.Stop()
}
