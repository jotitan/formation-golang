package model

// Ack interface to receive status of task
type Ack interface {
	Do(id int, status string)
}

//AckManager is an Ack implementation
type AckManager struct {
}

func NewAckManager(url string) *AckManager {
	return nil
}

func (a AckManager) Do(id int, status string) {
}
