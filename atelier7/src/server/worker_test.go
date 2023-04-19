package server

import (
	"formation-go/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const defaultWorkerPort = 9008

type mockAck struct {
	results map[int]string
	chanel  chan statusAck
}

type statusAck struct {
	id     int
	uuid   string
	status string
}

func (ma mockAck) Do(id int, uuid, status string) {
	ma.results[id] = status
	ma.chanel <- statusAck{id, uuid, status}
}

func newMockAck() mockAck {
	return mockAck{make(map[int]string), make(chan statusAck, 10)}
}

type mockRegister struct{}

func (mr mockRegister) Register(url, uuid string) error {
	return nil
}

func startAndWaitWorkerServerWithRegister(port int, ackManager model.Ack, register RegisterCoordinator) (Worker, error) {
	server := NewWorker(port, ackManager, register, false)
	if err := startGenericServer(server, port); err != nil {
		return Worker{}, err
	}
	return server, nil
}

func startStopWorkerServerWithRegister(t *testing.T, port int, ackManager model.Ack, register RegisterCoordinator, fct func(t *testing.T)) {
	server, err := startAndWaitWorkerServerWithRegister(port, ackManager, register)
	assert.Nil(t, err, "Server coordinator must start and be up")

	fct(t)

	server.Stop()
}

func startStopWorkerServer(t *testing.T, ackManager model.Ack, fct func(t *testing.T)) {
	startStopWorkerServerWithRegister(t, defaultWorkerPort, ackManager, mockRegister{}, fct)
}

func TestRunWorkerServer(t *testing.T) {
	startStopWorkerServer(t, nil, func(t *testing.T) {})
}

func TestUnknownTask(t *testing.T) {
	ack := newMockAck()
	startStopWorkerServer(t, ack, func(t *testing.T) {
		// GIVEN
		payload := "{\"type\":\"unknown\"}"

		// WHEN
		resp, _ := http.Post("http://localhost:9008/tasks/1", "application/json", strings.NewReader(payload))

		// THEN
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestBadMethod(t *testing.T) {
	ack := newMockAck()
	startStopWorkerServer(t, ack, func(t *testing.T) {
		// WHEN
		resp, _ := http.Get("http://localhost:9008/tasks/1")

		// THEN
		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	})
}

func TestRunTask(t *testing.T) {
	ack := newMockAck()
	startStopWorkerServer(t, ack, func(t *testing.T) {
		// GIVEN
		payload := "{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"

		// WHEN
		resp, _ := http.Post("http://localhost:9008/tasks/1", "application/json", strings.NewReader(payload))

		// THEN
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		d := <-ack.chanel
		assert.Equal(t, 1, d.id)
		assert.Equal(t, "finish", d.status)
	})
}

func TestCompleteChain(t *testing.T) {
	manager := model.NewManager()
	pool := NewWorkerPool(nil)
	startStopCoordinatorServer(t, manager, pool, func(t *testing.T) {
		startStopWorkerServerWithRegister(t, defaultWorkerPort, model.NewAckManager("http://localhost:9007"), NewRegisterCoordinator("http://localhost:9007"), func(t *testing.T) {
			// GIVEN
			manager.Add(model.NewPrint("Vers l'infini et l'au dela", 1))
			payload := "{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"

			// WHEN
			resp, _ := http.Post("http://localhost:9008/tasks/1", "application/json", strings.NewReader(payload))

			// THEN
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			_, status := manager.GetWithStatus(1)
			assert.Equal(t, model.TaskFinish, status)
		})
	})
}

func TestFullCompleteChain(t *testing.T) {
	manager := model.NewManager()
	pool := NewWorkerPool(LaunchTask{})
	startStopCoordinatorServer(t, manager, pool, func(t *testing.T) {
		startStopWorkerServerWithRegister(t, defaultWorkerPort, model.NewAckManager("http://localhost:9007"), NewRegisterCoordinator("http://localhost:9007"), func(t *testing.T) {
			// GIVEN

			// WHEN
			payload := "{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"
			resp, _ := http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))

			// THEN
			assert.Equal(t, 1, pool.Size())
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
			_, status := manager.GetWithStatus(1)
			assert.Equal(t, model.TaskFinish, status)
		})
	})
}

func TestManyWorkers(t *testing.T) {
	manager := model.NewManager()
	pool := NewWorkerPool(LaunchTask{})
	startStopCoordinatorServer(t, manager, pool, func(t *testing.T) {
		startStopWorkerServerWithRegister(t, 9009, model.NewAckManager("http://localhost:9007"), NewRegisterCoordinator("http://localhost:9007"), func(t *testing.T) {
			startStopWorkerServerWithRegister(t, defaultWorkerPort, model.NewAckManager("http://localhost:9007"), NewRegisterCoordinator("http://localhost:9007"), func(t *testing.T) {
				// GIVEN

				// WHEN
				payload := "{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"
				http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))
				http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))
				http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))
				http.Post("http://localhost:9007/tasks", "application/json", strings.NewReader(payload))

				// THEN
				assert.Equal(t, 2, pool.Size())
				for _, task := range manager.GetAllWithStatus() {
					assert.Equal(t, model.TaskFinish, task.Status)
				}
			})
		})
	})
}
