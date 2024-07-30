package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type TodoTask interface {
	CreateTask(user todo.Task) (string, error)
	UpdateTask(user todo.Task, id string) (int, error)
	DeleteTask(id string) (int, error)
	UpdateTaskAsDone(id string) (int, error)
	GetTask() ([]todo.Task, error)
	GetDoneTask() ([]todo.Task, error)
}

type Repository struct {
	TodoTask TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoTask: NewTaskPostgres(db),
	}
}
