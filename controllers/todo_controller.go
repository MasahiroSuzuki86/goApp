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
	UserId  string `json:"userid" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type TodoSearchRequest struct {
	UserId string `json:"userid" binding:"required"`
}

type TodoUpdateRequest struct {
	TodoId  string `json:"todoid" binding:"required"`
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

func (ctrl *TodoController) SearchTodo(c *gin.Context) {
	var req TodoSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	todos, err := ctrl.Service.SearchTodo(req.UserId)
	if err != nil {
		log.Println("Error retrieving todos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(todos) == 0 {
		log.Println("No todos found for user_id:", req.UserId)
		c.JSON(http.StatusNotFound, gin.H{"message": "No todos found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (ctrl *TodoController) UpdateTodo(c *gin.Context) {
	var req TodoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updateErr := ctrl.Service.UpdateTodo(req.TodoId, req.Content)
	if updateErr != nil {
		log.Println("Error todo update failed:", updateErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todo update err"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo update successfully"})
}
