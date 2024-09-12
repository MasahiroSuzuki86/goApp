package routes

import (
	"goApp/controllers"
	"goApp/repository"
	"goApp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userController.RegisterUser)
		userGroup.POST("/login", userController.LoginUser)
	}
}
