package server

import (
	"formation-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWorker(t *testing.T) {
	// GIVEN
	pool := NewWorkerPool(NewMockSender())

	// WHEN
	pool.Register("url1", "uuid1")

	// THEN
	assert.Equal(t, 1, pool.Size())
}

func TestAddWorkers(t *testing.T) {
	// GIVEN
	pool := NewWorkerPool(NewMockSender())

	// WHEN
	pool.Register("url1", "uuid1")
	pool.Register("url2", "uuid2")
	pool.Register("url3", "uuid3")

	// THEN
	assert.Equal(t, 3, pool.Size())
	assert.Equal(t, "uuid2", pool.workers["uuid2"].uuid)
	assert.Equal(t, "url2", pool.workers["uuid2"].url)
}

func TestAvoidDuplicate(t *testing.T) {
	// GIVEN
	pool := NewWorkerPool(NewMockSender())

	// WHEN
	pool.Register("url1", "uuid1")
	assert.False(t, pool.Register("url1", "uuid1"))

	// THEN
	assert.Equal(t, 1, pool.Size())
}

func TestSendTask(t *testing.T) {
	// GIVEN
	sender := NewMockSender()
	pool := NewWorkerPool(sender)
	pool.Register("url1", "uuid1")

	// WHEN
	pool.Execute(model.NewPrint("Vers l'infini et au dela", 1))

	// THEN
	assert.Equal(t, "url1", sender.store[1])
}

func TestSendTasksRoundRobin(t *testing.T) {
	// GIVEN
	sender := NewMockSender()
	pool := NewWorkerPool(sender)
	pool.Register("url1", "uuid1")
	pool.Register("url2", "uuid2")
	pool.Register("url3", "uuid3")

	// WHEN
	pool.Execute(model.NewPrint("Vers l'infini et au dela", 1))
	pool.Execute(model.NewPrint("Bonjour mon ami", 2))
	pool.Execute(model.NewPrint("Ou est mon oreille", 3))
	pool.Execute(model.NewPrint("Je sais voler", 4))

	// THEN
	assert.Equal(t, "url1", sender.store[1])
	assert.Equal(t, "url2", sender.store[2])
	assert.Equal(t, "url3", sender.store[3])
	assert.Equal(t, "url1", sender.store[4])
}

type MockSender struct {
	// Store, for task id, url where task is sended
	store map[int]string
}

func NewMockSender() MockSender {
	return MockSender{make(map[int]string)}
}

func (ms MockSender) Send(task model.Task, url string) error {
	ms.store[task.Id()] = url
	return nil
}
