package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"formation-go/model"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func startAndWaitCoordinatorServer(manager *model.Manager) (Coordinator, error) {
	server := NewCoordinator(9007, manager)
	if err := startGenericServer(server, 9007); err != nil {
		return Coordinator{}, err
	}
	return server, nil
}

func startStopCoordinatorServer(t *testing.T, manager *model.Manager, fct func(t *testing.T)) {
	server, err := startAndWaitCoordinatorServer(manager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	fct(t)

	server.Stop()
}

func TestRunServer(t *testing.T) {
	startStopCoordinatorServer(t, nil, func(t *testing.T) {})
}

func TestStatus(t *testing.T) {
	startStopCoordinatorServer(t, nil, func(t *testing.T) {
		// WHEN
		resp, _ := http.Get("http://localhost:9007/status")

		// THEN
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestStatusOnlyGet(t *testing.T) {
	startStopCoordinatorServer(t, nil, func(t *testing.T) {
		// WHEN
		resp, _ := http.Post("http://localhost:9007/status", "application/json", bytes.NewBuffer([]byte{}))

		// THEN
		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	})
}

func TestPostTaskPrint(t *testing.T) {
	manager := model.NewManager()
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		// WHEN
		resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))

		// THEN
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		idTask := string(read(resp.Body))
		assert.Equal(t, fmt.Sprintf("%d", manager.GetAll()[0].Id()), idTask)
		assert.Equal(t, 1, manager.Size())
	})
}

func TestUpdateTaskStatus(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	manager.Add(model.NewPrint("Vers l'infini et l'au dela", 1))

	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		// WHEN
		_, err := http.Post("http://localhost:9007/tasks/1", "application/json", strings.NewReader("{\"status\":\"finish\"}"))
		resp, _ := http.Get("http://localhost:9007/tasks/1")

		// THEN
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var task taskDto
		err = json.Unmarshal(read(resp.Body), &task)
		assert.Nil(t, err)
		assert.Equal(t, "finish", task.Status)
	})
}

func TestPostTaskMultiPrint(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		// WHEN
		_, err := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))
		_, err = http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"))
		_, err = http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Je sais voler\"}"))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, 3, manager.Size())
	})

}

func TestGetDetailPrintTask(t *testing.T) {
	// GIVEN
	startStopCoordinatorServer(t, model.NewManager(), func(t *testing.T) {
		resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"print\",\"message\":\"Bonjour mon ami\"}"))
		idTask := string(read(resp.Body))

		// WHEN
		resp, err := http.Get(fmt.Sprintf("http://localhost:9007/tasks/%s", idTask))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response map[string]interface{}
		err = json.Unmarshal(read(resp.Body), &response)

		assert.Nil(t, err)
		assert.Equal(t, "print", response["type"])
		assert.Equal(t, idTask, fmt.Sprintf("%.0f", response["id"]))
		assert.Equal(t, "running", response["status"])
	})
}

func read(reader io.Reader) []byte {
	data, _ := ioutil.ReadAll(reader)
	return data
}

func TestGetMultiTask(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		manager.Add(model.NewPrint("bonjour mon ami", 1))
		manager.Add(model.NewPrint("je sais voler", 2))
		manager.Add(model.NewPrint("c'est l'anniversaire d'andy", 3))

		// WHEN
		resp, _ := http.Get("http://localhost:9007/tasks")

		// THEN
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var lightTasks []taskDto
		err := json.Unmarshal(read(resp.Body), &lightTasks)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(lightTasks))
	})
}

func TestPostTaskResize(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		payload := "{\"type\":\"resize\",\"path\":\"/file.png\",\"target\":\"/file.png\", \"height\":200, \"width\":300}"
		resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))

		// THEN
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, 1, manager.Size())
	})
}

func TestPostTaskFail(t *testing.T) {
	// GIVEN
	manager := model.NewManager()
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		// WHEN
		resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader("{\"type\":\"unknown\"}"))

		// THEN
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, 0, manager.Size())
	})
}
