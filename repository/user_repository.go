package repository

import (
	"goApp/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ユーザーをDBに追加
func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

// ユーザー名でユーザーを検索
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// ユーザーのパスワードを変更する
func (repo *UserRepository) UpdatePassword(username string, newPassword string) error {
	return repo.DB.Model(&models.User{}).Where("username = ?", username).Update("password", newPassword).Error
}

// ユーザーを削除する
func (repo *UserRepository) DeleteUser(username string) error {
	return repo.DB.Where("username = ?", username).Delete(&models.User{}).Error
}
