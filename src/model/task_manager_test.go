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
