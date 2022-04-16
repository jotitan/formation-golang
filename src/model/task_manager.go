package model

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Add(task Task)    {}
func (m *Manager) Remove(task Task) {}
func (m *Manager) GetAll() []Task {
	return nil
}

func (m *Manager) Size() int { return 0 }

//GroupByType return a map where key is a string type and list contains of task of this type
func (m *Manager) GroupByType() map[string][]Task {
	return nil
}
