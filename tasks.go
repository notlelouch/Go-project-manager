package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.HandleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.HandleGetTask).Methods("GET")

}

func (s *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}
