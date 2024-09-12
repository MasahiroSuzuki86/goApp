// services/user_service.go
package services

import (
	"fmt"
	"goApp/models"
	"goApp/repository"
	"log"

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
		log.Fatalln(err)
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 新しいユーザーを作成
	newUser := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// ユーザーをデータベースに保存
	if err := s.Repo.CreateUser(newUser); err != nil {
		log.Fatalln(err)
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return newUser, nil
}

// LoginUser ログインサービス
func (s *UserService) LoginUser(username, password string) (*models.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	return user, nil
}
