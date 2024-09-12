package main

import (
	"goApp/config"
	"goApp/repository"
	"goApp/routes"
	"goApp/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoggingSettings(&config.Config)
	db := repository.Connect(&config.Config)
	router := gin.Default()
	routes.RegisterUserRoutes(router, db)
	router.Run(":8080")
}
