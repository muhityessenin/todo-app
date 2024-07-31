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

// createTask godoc
// @Summary Create a new task
// @Description Create a new task with the provided details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body todo.TaskInput true "Task Details"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/todo-list/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var input todo.Task
	var response Response
	if err := c.BindJSON(&input); err != nil {
		response = newResponse(http.StatusBadRequest, "bad request", "")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if h.validator.ValidateTaskInput(input) == false {
		response = newResponse(http.StatusBadRequest, "invalid input", "")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id, err := h.services.TodoTask.CreateTask(input)
	if err != nil {
		response = newResponse(http.StatusNotFound, "task is already exists", "")
		c.JSON(http.StatusNotFound, response)
		return
	}
	response = newResponse(http.StatusCreated, "task is created", id)
	c.JSON(http.StatusCreated, response)
}

// updateTaskAsDone godoc
// @Summary Mark a task as done
// @Description Mark a task as done by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 {object} Response
// @Failure 404 {object} Response
// @Router /api/todo-list/tasks/{id}/done [put]
func (h *Handler) updateTaskAsDone(c *gin.Context) {
	_, err := h.services.TodoTask.UpdateTaskAsDone(c.Param("id"))
	var res Response
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res = newResponse(http.StatusNotFound, "task not found", "")
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = newResponse(http.StatusNoContent, "task is updated", "")
	c.JSON(http.StatusNoContent, res)
}

// getTask godoc
// @Summary Get tasks
// @Description Get all tasks or filter by status
// @Tags tasks
// @Produce json
// @Param status query string false "Filter by status" Enums(done)
// @Success 200 {array} todo.Task
// @Failure 400 {object} Response
// @Router /api/todo-list/tasks [get]
func (h *Handler) getTask(c *gin.Context) {
	status := c.Query("status")
	var res Response
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
		res = newResponse(http.StatusBadRequest, "bad request", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = newResponse(http.StatusOK, "tasks successfully found", resTasks)
	c.JSON(http.StatusOK, res)
}

// updateTask godoc
// @Summary Update a task
// @Description Update a task's details
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body todo.TaskInput true "Task object"
// @Success 204 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/todo-list/tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	var input todo.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	input.Status = "FALSE"
	var res Response
	if h.validator.ValidateTaskInput(input) == false {
		res = newResponse(http.StatusBadRequest, "invalid input", "")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := h.services.TodoTask.UpdateTask(input, c.Param("id"))
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res = newResponse(http.StatusNotFound, "task not found", "")
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = newResponse(http.StatusNoContent, "task is updated", "")
	c.JSON(http.StatusNoContent, res)
}

// deleteTaskById godoc
// @Summary Delete a task by ID
// @Description Delete a task using its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/todo-list/tasks/{id} [delete]
func (h *Handler) deleteTaskById(c *gin.Context) {
	id := c.Param("id")
	var res Response
	_, err := h.services.TodoTask.DeleteTask(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		res = newResponse(http.StatusNotFound, "task not found", "")
		c.JSON(http.StatusNotFound, res)
		return
	}
	res = newResponse(http.StatusOK, "task is deleted", "")
	c.JSON(http.StatusOK, res)
}
