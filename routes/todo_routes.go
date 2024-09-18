package routes

import (
	"goApp/controllers"
	"goApp/repository"
	"goApp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTodoRoutes(router *gin.Engine, db *gorm.DB) {
	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	todoGroup := router.Group("/todo")
	{
		todoGroup.POST("/register", todoController.RegisterTodo)
		todoGroup.POST("/search", todoController.SearchTodo)
		todoGroup.POST("/update", todoController.UpdateTodo)
	}
}
