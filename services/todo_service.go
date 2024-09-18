package services

import (
	"fmt"
	"goApp/models"
	"goApp/repository"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

// RegisterTodo Todo登録サービス
func (s *TodoService) RegisterTodo(user_id string, content string) (*models.Todo, error) {
	newTodo := &models.Todo{
		Content: content,
		UserID:  user_id,
		Done:    false,
	}

	//todoをデータベースに保存
	if err := s.Repo.CreateTodo(newTodo); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to register todo: %w", err)
	}

	return newTodo, nil
}

// ユーザーIDでtodo一覧を取得する
func (s *TodoService) SearchTodo(userId string) ([]models.Todo, error) {
	todos, err := s.Repo.FindByUserId(userId)
	if err != nil {
		log.Println(err)
	} else if len(todos) == 0 {
		return nil, fmt.Errorf("No todos found")
	}

	return todos, nil
}

// todoを更新する
func (s *TodoService) UpdateTodo(todoId, content string) error {

	todoIdUint, err := strconv.ParseUint(todoId, 10, 32)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return err
	}

	updateTodo := &models.Todo{
		Model: gorm.Model{
			ID: uint(todoIdUint),
		},
		Content: content,
	}

	DBErr := s.Repo.UpdateTodo(updateTodo)
	if DBErr != nil {
		log.Println("Update failed:", DBErr)
		return DBErr
	}

	return nil
}
