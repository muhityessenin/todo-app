package validator

import (
	"time"
	todo "todo-app"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateTaskInput(task todo.Task) bool {
	if task.Title == "" {
		return false
	}
	t, err := time.Parse("2006-01-02", task.ActiveAt)
	if t.Before(time.Now()) {
		return false
	}
	if err != nil {
		return false
	}
	validStatuses := map[string]bool{"Pending": true, "In Progress": true, "Completed": true, "": true, "FALSE": true}
	if !validStatuses[task.Status] {
		return false
	}
	return true
}
