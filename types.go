package main

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToID int64     `json:"assignedTo"`
	CreartedAt   time.Time `json:"createdAt"`
}

type User struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Password   string    `json:"password"`
	CreartedAt time.Time `json:"createdAt"`
}
