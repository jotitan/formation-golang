package model

import (
	"errors"
	"reflect"
)

type RunningTask struct {
	Task   Task
	Status string
}

type Manager struct {
	tasks   map[int]*RunningTask
	counter int
}

func NewManager() *Manager {
	return &Manager{make(map[int]*RunningTask, 0), 0}
}

func (m *Manager) NextId() int {
	m.counter++
	return m.counter
}

func (m *Manager) GetWithStatus(id int) (Task, string) {
	task, exist := m.tasks[id]
	if !exist {
		return nil, ""
	}
	return task.Task, task.Status
}

func (m *Manager) UpdateStatus(id int, status string) {
	task, exist := m.tasks[id]
	if exist {
		task.Status = status
	}
}

func (m *Manager) Add(task Task) Task {
	m.tasks[task.Id()] = &RunningTask{Task: task, Status: "running"}
	return task
}
func (m *Manager) Remove(task Task) {
	delete(m.tasks, task.Id())
}

func (m *Manager) GetAllWithStatus() []RunningTask {
	tasks := make([]RunningTask, 0, len(m.tasks))
	for _, t := range m.tasks {
		tasks = append(tasks, *t)
	}
	return tasks
}

func (m *Manager) GetAll() []Task {
	tasks := make([]Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		tasks = append(tasks, t.Task)
	}
	return tasks
}

func (m *Manager) Size() int {
	return len(m.tasks)
}

//GroupByType return a map where key is a string type and list contains of Task of this type
func (m *Manager) GroupByType() map[string][]Task {
	groups := make(map[string][]Task)
	for _, t := range m.tasks {
		typ := reflect.TypeOf(t.Task).Name()
		list, exist := groups[typ]
		if !exist {
			list = make([]Task, 0)
		}
		groups[typ] = append(list, t.Task)
	}
	return groups
}

func (m *Manager) DetectAndCreateTask(payload map[string]interface{}) (Task, error) {
	return DetectAndCreateTask(payload, func() int { return m.NextId() })
}

func DetectAndCreateTask(payload map[string]interface{}, nextId func() int) (Task, error) {
	switch payload["type"].(string) {
	case "print":
		return NewPrint(payload["message"].(string), nextId()), nil
	case "resize":
		return NewResize(payload["path"].(string), payload["target"].(string), int(payload["height"].(float64)), int(payload["width"].(float64)), nextId()), nil
	default:
		return nil, errors.New("unknown type")
	}
}
