package repository

import (
	"goApp/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// TodoをDBに追加
func (repo *TodoRepository) CreateTodo(todo *models.Todo) (err error) {
	return repo.DB.Create(todo).Error
}

// ユーザーIDでtodoを検索
func (repo *TodoRepository) FindByUserId(userId string) ([]models.Todo, error) {
	var todos []models.Todo
	err := repo.DB.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}
