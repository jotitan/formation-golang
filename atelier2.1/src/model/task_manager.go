package model

import (
	"reflect"
)

type Manager struct {
	tasks []Task
}

func NewManager() *Manager {
	return &Manager{make([]Task, 0)}
}

func (m *Manager) NextId() int {
	// Write code here
	return 0
}

func (m *Manager) Add(task Task) Task {
	m.tasks = append(m.tasks, task)
	return task
}
func (m *Manager) Remove(task Task) {
	// write code here
}

func (m *Manager) GetAll() []Task {
	return m.tasks
}

func (m *Manager) Size() int {
	return len(m.tasks)
}

//GroupByType return a map where key is a string type and list contains of task of this type
func (m *Manager) GroupByType() map[string][]Task {
	groups := make(map[string][]Task)
	for _, t := range m.tasks {
		typ := reflect.TypeOf(t).Name()
		list, exist := groups[typ]
		if !exist {
			list = make([]Task, 0)
		}
		groups[typ] = append(list, t)
	}
	return groups
}
