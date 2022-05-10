package model

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Add(task Task) {
	// Write code here
}

func (m *Manager) GetAll() []Task {
	// Write code here
	return nil
}

func (m *Manager) Size() int {
	// Write code here
	return 0
}

//GroupByType return a map where key is a string type and list contains of task of this type
func (m *Manager) GroupByType() map[string][]Task {
	// Write code here
	return nil
}
