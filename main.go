package main

import (
	"goApp/config"
	"goApp/repository"
	"goApp/routes"
	"goApp/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoggingSettings("webapp")
	db := repository.Connect(&config.Config)
	router := gin.Default()
	routes.RegisterUserRoutes(router, db)
	routes.RegisterTodoRoutes(router, db)
	router.Run(":8080")
}
