package model

import (
	"github.com/stretchr/testify/assert"
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
	manager.Add(NewPrint("", 0))
	// WHEN
	size := manager.Size()
	// THEN
	assert.Equal(t, 2, size, "Must find 2")
}

func TestGroupManager(t *testing.T) {
	// GIVEN
	manager := NewManager()
	manager.Add(NewPrint("", 0))
	manager.Add(NewPrint("", 0))
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
	manager.Add(NewPrint("", 0))
	manager.Add(NewResize("", "", 0, 0, 0))
	// WHEN
	groups := manager.GroupByType()
	// THEN
	assert.Equal(t, 2, len(groups), "Must find 2")
}

func TestRemoveMissingElement(t *testing.T) {
	manager := NewManager()
	task1 := NewPrint("", manager.NextId())
	task2 := NewPrint("", manager.NextId())

	if nb := manager.Size(); nb != 0 {
		t.Error("Must find 0, but find", nb)
	}

	manager.Add(task1)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1, but find", nb)
	}

	manager.Remove(task2)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1, but find", nb)
	}
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

	if nb := manager.Size(); nb != 3 {
		t.Error("Must find 3, but find", nb)
	}

	manager.Remove(task1)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2, but find", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must find", task2.Id())
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
		t.Error("Must find 4, but find", nb)
	}

	manager.Remove(task1)
	manager.Remove(task3)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2, but find", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must find", task2.Id())
	}
	if manager.GetAll()[1].Id() != task4.Id() {
		t.Error("Must find", task4.Id())
	}
}
