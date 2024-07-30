package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type TodoTask interface {
	CreateTask(task todo.Task) (string, error)
	UpdateTask(task todo.Task, id string) (int, error)
	DeleteTask(id string) (int, error)
	UpdateTaskAsDone(id string) (int, error)
	GetTask() ([]todo.Task, error)
	GetDoneTask() ([]todo.Task, error)
}

type Service struct {
	TodoTask TodoTask
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTaskService(repository.TodoTask),
	}
}
