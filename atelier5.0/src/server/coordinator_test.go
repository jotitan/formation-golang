package server

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func startAndWaitServer() (Coordinator, error) {
	server := NewCoordinator(9007)
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
	server, error := startAndWaitServer()

	// WHEN - THEN
	assert.Nil(t, error, "Server coordinator must start and be up")
	server.Stop()
}

func TestStatusOnlyGet(t *testing.T) {
	// GIVEN
	server, error := startAndWaitServer()
	assert.Nil(t, error, "Server coordinator must start and be up")

	// WHEN
	resp, _ := http.Post("http://localhost:9007/status", "application/json", bytes.NewBuffer([]byte{}))

	// THEN
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	server.Stop()
}
