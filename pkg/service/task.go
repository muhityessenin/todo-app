package service

import (
	"fmt"
	"time"
	todo "todo-app"
	"todo-app/pkg/repository"
)

type TaskService struct {
	repo repository.TodoTask
}

func NewTaskService(repo repository.TodoTask) *TaskService {
	return &TaskService{repo: repo}
}

func (t TaskService) CreateTask(task todo.Task) (string, error) {
	return t.repo.CreateTask(task)
}

func (t TaskService) UpdateTask(task todo.Task, id string) (int, error) {
	return t.repo.UpdateTask(task, id)
}

func (t TaskService) DeleteTask(id string) (int, error) {
	return t.repo.DeleteTask(id)
}

func (t TaskService) UpdateTaskAsDone(id string) (int, error) {
	return t.repo.UpdateTaskAsDone(id)
}

func (t TaskService) GetTask() ([]todo.Task, error) {
	tasks, err := t.repo.GetTask()
	if err != nil {
		return nil, err
	}
	var res []todo.Task

	for _, task := range tasks {
		activeAt, err := time.Parse("2006-01-02", task.ActiveAt[:10])
		if err != nil {
			fmt.Println(err)
		}
		if activeAt.Weekday().String() == "Sunday" || activeAt.Weekday().String() == "Saturday" {
			task.Title = "ВЫХОДНОЙ - " + task.Title
		}
		if activeAt.After(time.Now()) || activeAt.Equal(time.Now()) {
			task.ActiveAt = task.ActiveAt[:10]
			res = append(res, task)
		}
	}
	return res, nil
}

func (t TaskService) GetDoneTask() ([]todo.Task, error) {
	tasks, err := t.repo.GetDoneTask()
	if err != nil {
		return nil, err
	}
	var res []todo.Task
	for _, task := range tasks {
		activeAt, err := time.Parse("2006-01-02", task.ActiveAt[:10])
		if err != nil {
			fmt.Println(err)
		}
		if activeAt.Weekday().String() == "Sunday" || activeAt.Weekday().String() == "Saturday" {
			task.Title = "ВЫХОДНОЙ - " + task.Title
		}
		if activeAt.After(time.Now()) || activeAt.Equal(time.Now()) {
			task.ActiveAt = task.ActiveAt[:10]
			res = append(res, task)
		}
	}
	return res, nil
}
