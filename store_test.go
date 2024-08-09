package main

type MockStore struct{}

func (m *MockStore) CreateUser() error {
	return nil
}

func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetTask() (*Task, error) {
	return &Task{}, nil
}