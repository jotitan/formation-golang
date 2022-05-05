package model

import "testing"

func TestManager(t *testing.T) {
	manager := NewManager()
	if nb := manager.Size(); nb != 0 {
		t.Error("Must find 0 but found", nb)
	}

	manager.Add(NewPrint("", 0))
	if nb := manager.Size(); nb != 1 {
		t.Error("Must find 1 but found", nb)
	}
	manager.Add(NewPrint("", 0))
	if nb := manager.Size(); nb != 2 {
		t.Error("Must find 2 but found", nb)
	}

	groups := manager.GroupByType()
	if nb := len(groups); nb != 1 {
		t.Error("Must find 1 but found", nb)
	}
	if nb := len(groups["print"]); nb != 2 {
		t.Error("Must find 2 but found", nb)
	}
	manager.Add(NewResize("", "", 0, 0, 0))
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
