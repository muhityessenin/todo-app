package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo-app/pkg/service"
	"todo-app/pkg/validator"
)

type Handler struct {
	services  *service.Service
	validator *validator.Validator
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		tasks := api.Group("/todo-list")
		{
			tasks.POST("/tasks", h.createTask)
			tasks.PUT("/tasks/:id", h.updateTask)
			tasks.PUT("/tasks/:id/done", h.updateTaskAsDone)
			tasks.GET("/tasks", h.getTask)
			tasks.DELETE("/tasks/:id", h.deleteTaskById)
		}
	}
	return router
}
