package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	todo "todo-app"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) isTaskPresent(title string, activeAt string) bool {
	var task todo.Task
	query := fmt.Sprintf("SELECT id FROM %s WHERE title=$1 AND active_at=$2", "tasks")
	err := r.db.QueryRow(query, title, activeAt).Scan(&task.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		log.Fatalf("Error checking task existence: %v", err)
	}
	return true
}

func (r *TaskPostgres) Check(id string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", TasksTable)
	err := r.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, sql.ErrNoRows
		}
		return false, fmt.Errorf("error checking task existence: %v", err)
	}
	return true, nil
}

func (r *TaskPostgres) CreateTask(task todo.Task) (string, error) {
	var id string
	check := r.isTaskPresent(task.Title, task.ActiveAt)
	if check {
		return "", errors.New("task is already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (title, active_at) values ($1, $2) RETURNING id", TasksTable)
	row := r.db.QueryRow(query, task.Title, task.ActiveAt)
	if err := row.Scan(&id); err != nil {
		return strconv.Itoa(0), err
	}
	return id, nil
}

func (r *TaskPostgres) UpdateTask(task todo.Task, id string) (int, error) {
	_, err := r.Check(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, sql.ErrNoRows
		}
		return http.StatusInternalServerError, err
	}
	query := fmt.Sprintf("UPDATE %s SET title = $1, active_at = $2, status = $3 WHERE id = $4", TasksTable)
	_, err = r.db.Exec(query, task.Title, task.ActiveAt, task.Status, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (r *TaskPostgres) DeleteTask(id string) (int, error) {
	_, err := r.Check(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, sql.ErrNoRows
		}
		return http.StatusInternalServerError, err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", TasksTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (r *TaskPostgres) UpdateTaskAsDone(id string) (int, error) {
	_, err := r.Check(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, sql.ErrNoRows
		}
		return http.StatusInternalServerError, err
	}
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", TasksTable)
	_, err = r.db.Exec(query, "TRUE", id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (r *TaskPostgres) GetTask() ([]todo.Task, error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY active_at, status", TasksTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return make([]todo.Task, 0), err
	}
	defer rows.Close()

	var tasks []todo.Task
	for rows.Next() {
		var task todo.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.ActiveAt, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return make([]todo.Task, 0), err
	}

	return tasks, nil
}
func (r *TaskPostgres) GetDoneTask() ([]todo.Task, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE status = TRUE ORDER BY active_at", TasksTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return make([]todo.Task, 0), err
	}
	defer rows.Close()

	var tasks []todo.Task
	for rows.Next() {
		var task todo.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.ActiveAt, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return make([]todo.Task, 0), err
	}
	return tasks, nil
}
