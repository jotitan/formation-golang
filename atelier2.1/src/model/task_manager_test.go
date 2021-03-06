package model

import (
	"testing"
)

func TestManager(t *testing.T) {
	manager := NewManager()
	if nb := manager.Size(); nb != 0 {
		t.Error("Must find 0 but found", nb)
	}

	manager.Add(Print{})
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1 but found", nb)
	}
	manager.Add(Print{})
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2 but found", nb)
	}

	groups := manager.GroupByType()
	if nb := len(groups); nb != 1 {
		t.Error("Must find 1 but found", nb)
	}
	if nb := len(groups["Print"]); nb != 2 {
		t.Error("Must find 2 but found", nb)
	}
	manager.Add(Resize{})
	if nb := manager.Size(); nb != 3 {
		t.Error("Must find 3 but found", nb)
	}

	groups = manager.GroupByType()
	if nb := len(groups); nb != 2 {
		t.Error("Must find 2 but found", nb)
	}
}

func TestRemoveMissingElement(t *testing.T) {
	manager := NewManager()
	task1 := Print{Uuid: manager.NextId()}
	task2 := Print{Uuid: manager.NextId()}

	if nb := manager.Size(); nb != 0 {
		t.Error("Must find 0, but found", nb)
	}

	manager.Add(task1)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1, but found", nb)
	}

	manager.Remove(task2)
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1, but found", nb)
	}
}

func TestSimpleRemove(t *testing.T) {
	manager := NewManager()
	task1 := Print{Uuid: manager.NextId()}
	task2 := Print{Uuid: manager.NextId()}
	task3 := Print{Uuid: manager.NextId()}

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)

	if nb := manager.Size(); nb != 3 {
		t.Error("Must find 3, but found", nb)
	}

	manager.Remove(task1)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2, but found", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must find", task2.Id())
	}
}

func TestComplexRemove(t *testing.T) {
	manager := NewManager()
	task1 := Print{Uuid: manager.NextId()}
	task2 := Print{Uuid: manager.NextId()}
	task3 := Print{Uuid: manager.NextId()}
	task4 := Print{Uuid: manager.NextId()}

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)
	manager.Add(task4)

	if nb := manager.Size(); nb != 4 {
		t.Error("Must find 4, but found", nb)
	}

	manager.Remove(task1)
	manager.Remove(task3)
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2, but found", nb)
	}
	if manager.GetAll()[0].Id() != task2.Id() {
		t.Error("Must find", task2.Id())
	}
	if manager.GetAll()[1].Id() != task4.Id() {
		t.Error("Must find", task4.Id())
	}
}
