package main

import (
	"goApp/repository"
	"goApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.Connect()
	router := gin.Default()
	routes.RegisterUserRoutes(router, db)

	router.Run(":8080")
}
