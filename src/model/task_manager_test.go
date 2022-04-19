package model

import "testing"
import "github.com/stretchr/testify/assert"

func TestManager(t *testing.T) {
	manager := NewManager()
	assert.Equal(t, 0, manager.Size())

	manager.Add(Print{Id: manager.NextId()})
	assert.Equal(t, 1, manager.Size())
	manager.Add(Print{Id: manager.NextId()})
	assert.Equal(t, 2, manager.Size())

	groups := manager.GroupByType()
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, 2, len(groups["Print"]))

	manager.Add(Resize{Id: manager.NextId()})
	assert.Equal(t, 3, manager.Size())

	groups = manager.GroupByType()
	assert.Equal(t, 2, len(groups))
	assert.Equal(t, 1, len(groups["Resize"]))
}

func TestRemoveMissingElement(t *testing.T) {
	manager := NewManager()
	task1 := Print{Id: manager.NextId()}
	task2 := Print{Id: manager.NextId()}

	assert.Equal(t, 0, manager.Size())

	manager.Add(task1)
	assert.Equal(t, 1, manager.Size())

	manager.Remove(task2)
	assert.Equal(t, 1, manager.Size())
}

func TestSimpleRemove(t *testing.T) {
	manager := NewManager()
	task1 := Print{Id: manager.NextId()}
	task2 := Print{Id: manager.NextId()}
	task3 := Print{Id: manager.NextId()}

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)

	assert.Equal(t, 3, manager.Size())

	manager.Remove(task1)
	assert.Equal(t, 2, manager.Size())
	assert.Equal(t, task2.GetId(), manager.GetAll()[0].GetId())
}

func TestComplexRemove(t *testing.T) {
	manager := NewManager()
	task1 := Print{Id: manager.NextId()}
	task2 := Print{Id: manager.NextId()}
	task3 := Print{Id: manager.NextId()}
	task4 := Print{Id: manager.NextId()}

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)
	manager.Add(task4)

	assert.Equal(t, 4, manager.Size())

	manager.Remove(task1)
	manager.Remove(task3)
	assert.Equal(t, 2, manager.Size())
	assert.Equal(t, task2.GetId(), manager.GetAll()[0].GetId())
	assert.Equal(t, task4.GetId(), manager.GetAll()[1].GetId())
}