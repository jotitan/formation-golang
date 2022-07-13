package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type serverToStart interface {
	Run()
}

func startGenericServer(server serverToStart, port int) error {
	go server.Run()
	//log.Println("Run server")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		// Call api status
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d/status", port))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
	}
	return errors.New("impossible to start server")
}
