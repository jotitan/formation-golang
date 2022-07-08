package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Coordinator struct {
	server *http.Server
}

func NewCoordinator(port int) Coordinator {
	server := http.ServeMux{}
	server.HandleFunc("/status", status)

	wrapServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &server,
	}

	return Coordinator{&wrapServer}
}

func (c Coordinator) Run() {
	c.server.ListenAndServe()
}

func (c Coordinator) Stop() {
	c.server.Shutdown(context.Background())
}

func status(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, http.MethodGet) {
		http.Error(w, "Not allowed method", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("up"))
}
