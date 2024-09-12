// services/user_service.go
package services

import (
	"fmt"
	"goApp/models"
	"goApp/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService ユーザー関連のサービスを定義
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService サービスの初期化
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// RegisterUser ユーザー登録サービス
func (s *UserService) RegisterUser(username, password string) (*models.User, error) {
	// ユーザー名の重複チェック
	existingUser, err := s.Repo.FindByUsername(username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 新しいユーザーを作成
	newUser := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// ユーザーをデータベースに保存
	if err := s.Repo.CreateUser(newUser); err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return newUser, nil
}
