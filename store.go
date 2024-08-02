package main

import (
	"database/sql"
)

type Store interface {
	// users
	CreateUser() error

	// tasks
	CreateTask(task *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id,assigned_to) VALUES (?, ?, ?, ?)",
		task.Name, task.Status, task.ProjectId, task.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	task.ID = id
	return task, nil
}
