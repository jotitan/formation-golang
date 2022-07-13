package server

import (
	"formation-go/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

type mockAck struct {
	results map[int]string
	chanel  chan statusAck
}

type statusAck struct {
	id     int
	status string
}

func (ma mockAck) Do(id int, status string) {
	ma.results[id] = status
	ma.chanel <- statusAck{id, status}
}

func newMockAck() mockAck {
	return mockAck{make(map[int]string), make(chan statusAck, 10)}
}

func startAndWaitWorkerServer(ackManager model.Ack) (Worker, error) {
	server := NewWorker(9008, ackManager, false)
	if err := startGenericServer(server, 9008); err != nil {
		return Worker{}, err
	}
	return server, nil
}

func startStopWorkerServer(t *testing.T, ackManager model.Ack, fct func(t *testing.T)) {
	server, err := startAndWaitWorkerServer(ackManager)
	assert.Nil(t, err, "Server coordinator must start and be up")

	fct(t)

	server.Stop()
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
	startStopCoordinatorServer(t, manager, func(t *testing.T) {
		startStopWorkerServer(t, model.NewAckManager("http://localhost:9007"), func(t *testing.T) {
			// GIVEN
			manager.Add(model.NewPrint("Vers l'infini et l'au dela", 1))
			payload := "{\"type\":\"print\",\"message\":\"Vers l'infini et l'au dela\"}"

			// WHEN
			resp, _ := http.Post("http://localhost:9008/tasks/1", "application/json", strings.NewReader(payload))

			// THEN
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			_, status := manager.GetWithStatus(1)
			assert.Equal(t, "finish", status)
		})
	})
}
