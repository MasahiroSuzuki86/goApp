package controllers

import (
	"goApp/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{Service: service}
}

type TodoRegisterRequest struct {
	UserId  uint   `json:"userid" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (ctrl *TodoController) RegisterTodo(c *gin.Context) {
	var req TodoRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "INvalid request"})
		return
	}

	_, err := ctrl.Service.RegisterTodo(req.UserId, req.Content)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo registered successfully"})
}
