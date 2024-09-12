package controllers

import (
	"goApp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 構造体
type UserController struct {
	Service *services.UserService
}

// NewUserController コントローラーの初期化
func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// RegisterRequest リクエストボディの構造体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser ユーザー登録のエンドポイント
func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// サービスを使ってユーザーを登録
	_, err := ctrl.Service.RegisterUser(req.Username, req.Password)
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
