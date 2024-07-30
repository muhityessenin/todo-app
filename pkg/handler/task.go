package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	todo "todo-app"
)

type ResponseTask struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt" date_format:"YYYY-MM-DD"`
}

func (h *Handler) createTask(c *gin.Context) {
	var input todo.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	id, err := h.services.TodoTask.CreateTask(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task is already created",
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) updateTaskAsDone(c *gin.Context) {
	_, err := h.services.TodoTask.UpdateTaskAsDone(c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task is not found",
		})
	}
	c.JSON(http.StatusNoContent, nil)
}
func (h *Handler) getTask(c *gin.Context) {
	status := c.Query("status")

	var tasks []todo.Task
	var err error
	if status == "" {
		tasks, err = h.services.TodoTask.GetTask()
	} else if status == "done" {
		tasks, err = h.services.TodoTask.GetDoneTask()
	}
	resTasks := make([]ResponseTask, len(tasks))
	for i, task := range tasks {
		activeAt := task.ActiveAt
		if len(activeAt) > 10 {
			activeAt = activeAt[:10]
		}
		resTasks[i] = ResponseTask{
			ID:       task.ID,
			Title:    task.Title,
			ActiveAt: activeAt,
		}
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resTasks)
}
func (h *Handler) updateTask(c *gin.Context) {
	var input todo.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	input.Status = "FALSE"
	if h.validator.ValidateTaskInput(input) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	_, err := h.services.TodoTask.UpdateTask(input, c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func (h *Handler) deleteTaskById(c *gin.Context) {
	id := c.Param("id")
	_, err := h.services.TodoTask.DeleteTask(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
