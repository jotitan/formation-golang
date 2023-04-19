package model

import (
	"fmt"
	"net/http"
	"strings"
)

type Ack interface {
	Do(id int, uuid, status string)
}

type AckManager struct {
	urlCoordinator string
}

func NewAckManager(url string) *AckManager {
	return &AckManager{url}
}

func (a AckManager) Do(id int, uuid, status string) {
	url := fmt.Sprintf("%s/tasks/%d", a.urlCoordinator, id)
	http.Post(url, "application/json", strings.NewReader(fmt.Sprintf("{\"status\":\"%s\",\"uuid\":\"%s\"}", status, uuid)))
}
