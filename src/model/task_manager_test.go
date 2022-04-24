package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager(t *testing.T) {
	manager := NewManager()
	assert.Equal(t, 0, manager.Size(), "Must found 0")

	manager.Add(NewPrint("", 0))
	assert.Equal(t, 1, manager.Size(), "Must found 1")

	manager.Add(NewPrint("", 0))
	assert.Equal(t, 2, manager.Size(), "Must found 2")

	groups := manager.GroupByType()
	assert.Equal(t, 1, len(groups), "Must found 1")

	assert.Equal(t, 2, len(groups["print"]), "Must found 2")

	manager.Add(NewResize("", "", 0, 0, 0))
	assert.Equal(t, 3, manager.Size(), "Must found 3")

	groups = manager.GroupByType()
	assert.Equal(t, 2, len(groups), "Must found 2")
}

func TestRemoveMissingElement(t *testing.T) {
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())

	if nb := manager.Size(); nb != 0 {
		t.Error("Must found 0, but found", nb)
	}

	manager.Add(task1)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must found 1, but found", nb)
	}

	manager.Remove(task2)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must found 1, but found", nb)
	}
}

func TestSimpleRemove(t *testing.T) {
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())
	task3 := NewPrint("", manager.NextId())

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)

	if nb := manager.Size(); nb != 3 {
		t.Error("Must found 3, but found", nb)
	}

	manager.Remove(task1)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must found 2, but found", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must found", task2.Id())
	}
}

func TestComplexRemove(t *testing.T) {
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())
	task3 := NewPrint("", manager.NextId())
	task4 := NewPrint("", manager.NextId())

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)
	manager.Add(task4)

	if nb := manager.Size(); nb != 4 {
		t.Error("Must found 4, but found", nb)
	}

	manager.Remove(task1)
	manager.Remove(task3)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must found 2, but found", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must found", task2.Id())
	}
	if manager.GetAll()[1].Id() != task4.Id() {
		t.Error("Must found", task4.Id())
	}
}
