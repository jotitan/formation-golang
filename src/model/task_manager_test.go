package model

import "testing"

func TestManager(t *testing.T) {
	manager := NewManager()
	if nb := manager.Size(); nb != 0 {
		t.Error("Must found 0 but found", nb)
	}

	manager.Add(Print{})
	if nb := manager.Size(); nb != 1 {
		t.Error("Must found 1 but found", nb)
	}
	manager.Add(Print{})
	if nb := manager.Size(); nb != 2 {
		t.Error("Must found 2 but found", nb)
	}

	groups := manager.GroupByType()
	if nb := len(groups); nb != 1 {
		t.Error("Must found 1 but found", nb)
	}
	if nb := len(groups["Print"]); nb != 2 {
		t.Error("Must found 2 but found", nb)
	}
	manager.Add(Resize{})
	if nb := manager.Size(); nb != 3 {
		t.Error("Must found 3 but found", nb)
	}

	groups = manager.GroupByType()
	if nb := len(groups); nb != 2 {
		t.Error("Must found 2 but found", nb)
	}
}

func TestRemoveMissingElement(t *testing.T) {
	manager := NewManager()
	task1 := Print{Uuid: manager.NextId()}
	task2 := Print{Uuid: manager.NextId()}

	assert.Equal(t, 0, manager.Size())

	manager.Add(task1)
	assert.Equal(t, 1, manager.Size())

	manager.Remove(task2)
	assert.Equal(t, 1, manager.Size())
}

func TestSimpleRemove(t *testing.T) {
	manager := NewManager()
	task1 := Print{Uuid: manager.NextId()}
	task2 := Print{Uuid: manager.NextId()}
	task3 := Print{Uuid: manager.NextId()}

	manager.Add(task1)
	manager.Add(task2)
	manager.Add(task3)

	assert.Equal(t, 3, manager.Size())

	manager.Remove(task1)
	assert.Equal(t, 2, manager.Size())
	assert.Equal(t, task2.Id(), manager.GetAll()[0].Id())
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

	assert.Equal(t, 4, manager.Size())

	manager.Remove(task1)
	manager.Remove(task3)
	assert.Equal(t, 2, manager.Size())
	assert.Equal(t, task2.Id(), manager.GetAll()[0].Id())
	assert.Equal(t, task4.Id(), manager.GetAll()[1].Id())
}
