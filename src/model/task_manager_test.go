package model

import "testing"
import "github.com/stretchr/testify/assert"

func TestManager(t *testing.T) {
	manager := NewManager()
	assert.Equal(t, 0, manager.Size())

	manager.Add(Print{})
	assert.Equal(t, 1, manager.Size())
	manager.Add(Print{})
	assert.Equal(t, 2, manager.Size())

	groups := manager.GroupByType()
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, 2, len(groups["Print"]))

	manager.Add(Resize{})
	assert.Equal(t, 3, manager.Size())

	groups = manager.GroupByType()
	assert.Equal(t, 2, len(groups))
}
