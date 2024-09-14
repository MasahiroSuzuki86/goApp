package services

import (
	"fmt"
	"goApp/models"
	"goApp/repository"
	"log"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

// RegisterTodo Todo登録サービス
func (s *TodoService) RegisterTodo(user_id uint, content string) (*models.Todo, error) {
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
