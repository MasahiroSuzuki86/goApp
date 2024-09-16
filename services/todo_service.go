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

func (s *TodoService) SearchTodo(userId string) ([]models.Todo, error) {
	todos, err := s.Repo.FindByUserId(userId)
	if err != nil {
		log.Println(err)
	} else if len(todos) == 0 {
		return nil, fmt.Errorf("No todos found")
	}

	return todos, nil
}
