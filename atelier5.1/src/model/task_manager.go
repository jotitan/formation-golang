package model

import (
	"errors"
	"reflect"
)

type Manager struct {
	tasks   []Task
	counter int
}

func NewManager() *Manager {
	return &Manager{make([]Task, 0), 0}
}

func (m *Manager) NextId() int {
	m.counter++
	return m.counter
}

func (m *Manager) Get(id int) Task {
	for _, task := range m.tasks {
		if task.Id() == id {
			return task
		}
	}
	return nil
}

func (m *Manager) Add(task Task) Task {
	m.tasks = append(m.tasks, task)
	return task
}
func (m *Manager) Remove(task Task) {
	foundPosition := -1
	for pos, t := range m.tasks {
		if t.Id() == task.Id() {
			foundPosition = pos
			break
		}
	}
	switch foundPosition {
	case -1:
		break
	case 0:
		m.tasks = m.tasks[1:]
	case len(m.tasks) - 1:
		m.tasks = m.tasks[0 : len(m.tasks)-1]
	default:
		m.tasks = append(m.tasks[0:foundPosition], m.tasks[foundPosition+1:]...)
	}
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

func (m *Manager) DetectAndCreateTask(payload map[string]interface{}) (Task, error) {
	switch payload["type"].(string) {
	case "print":
		return NewPrint(payload["message"].(string), m.NextId()), nil
	case "resize":
		return NewResize(payload["path"].(string), payload["target"].(string), int(payload["height"].(float64)), int(payload["width"].(float64)), m.NextId()), nil
	default:
		return nil, errors.New("unknown type")
	}
}
