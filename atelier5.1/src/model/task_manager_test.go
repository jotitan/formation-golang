package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestEmptyManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	// WHEN
	size := manager.Size()
	// THEN
	assert.Equal(t, 0, size, "Must find 0")
}

func TestAddOneManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	manager.Add(NewPrint("", 0))
	// WHEN
	size := manager.Size()
	// THEN
	assert.Equal(t, 1, size, "Must find 1")
}

func TestAddTwoManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	manager.Add(NewPrint("", 0))
	manager.Add(NewPrint("", 1))
	// WHEN
	size := manager.Size()
	// THEN
	assert.Equal(t, 2, size, "Must find 2")
}

func TestGroupManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	manager.Add(NewPrint("", 0))
	manager.Add(NewPrint("", 1))
	// WHEN
	groups := manager.GroupByType()
	// THEN
	assert.Equal(t, 1, len(groups), "Must find 1")
	assert.Equal(t, 2, len(groups["print"]), "Must find 2")
}

func TestGroupTwoManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	manager.Add(NewPrint("", 0))
	manager.Add(NewPrint("", 1))
	manager.Add(NewResize("", "", 0, 0, 0))
	// WHEN
	groups := manager.GroupByType()
	// THEN
	assert.Equal(t, 2, len(groups), "Must find 2")
}

func TestRemoveMissingElement(t *testing.T) {
	// GIVEN
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())

	// WHEN
	manager.Add(task1)
	assert.Equal(t, 1, manager.Size(), "Must find 1")
	manager.Remove(task2)

	// THEN
	assert.Equal(t, 1, manager.Size(), "Must find 1")
}

func TestSimpleRemove(t *testing.T) {
	// GIVEN
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())
	task3 := NewPrint("", manager.NextId())

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)
	assert.Equal(t, 3, manager.Size(), "Must find 3")

	// WHEN
	manager.Remove(task1)

	// THEN
	assert.Equal(t, 2, manager.Size(), "Must find 2")
	assert.Equal(t, task2.Id(), manager.GetAll()[0].Id(), fmt.Sprintf("Must find %d", task2.Id()))
}

func TestGetNotFound(t *testing.T) {
	// GIVEN
	manager := NewManager()

	// WHEN
	task, _ := manager.GetWithStatus(1)

	// THEN
	assert.Nil(t, task)
}

func TestUpdateStatus(t *testing.T) {
	// GIVEN
	manager := NewManager()
	task := NewPrint("", manager.NextId())
	manager.Add(task)

	// WHEN
	manager.UpdateStatus(task.Id(), "finish")

	// THEN
	_, status := manager.GetWithStatus(task.Id())
	assert.Equal(t, "finish", status)
}

func TestComplexRemove(t *testing.T) {
	// GIVEN
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())
	task3 := NewPrint("", manager.NextId())
	task4 := NewPrint("", manager.NextId())

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)
	manager.Add(task4)

	assert.Equal(t, 4, manager.Size(), "Must find 4")

	// WHEN
	manager.Remove(task1)
	manager.Remove(task3)

	// THEN
	assert.Equal(t, 2, manager.Size(), "Must find 2")
	tasks := manager.GetAll()
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id() < tasks[j].Id()
	})
	assert.Equal(t, task2.Id(), tasks[0].Id(), fmt.Sprintf("Must find id %d", task2.Id()))
	assert.Equal(t, task4.Id(), tasks[1].Id(), fmt.Sprintf("Must find id %d", task4.Id()))
}
