package repository

import (
	"fmt"
	"goApp/config"
	"goApp/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config *config.ConfigList) *gorm.DB {
	// 環境変数からデータベース接続情報を取得
	DB_USER := config.DB_USER
	DB_PASSWORD := config.DB_PASSWORD
	DB_NAME := config.DB_NAME
	DB_HOST := config.DB_HOST
	DB_PORT := config.DB_PORT

	// DSN(Data Source Name)を作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database: ", err)
	}

	userErr := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Failed to AutoMigrate: ", userErr)
	}
	toDoErr := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Println("Failed to AutoMigrate: ", toDoErr)
	}

	return db
}
