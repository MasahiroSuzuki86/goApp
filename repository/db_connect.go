package repository

import (
	"fmt"
	"goApp/config"
	"goApp/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// 環境変数からデータベース接続情報を取得
	DB_USER := config.Config.DB_USER
	DB_PASSWORD := config.Config.DB_PASSWORD
	DB_NAME := config.Config.DB_NAME
	DB_HOST := config.Config.DB_HOST
	DB_PORT := config.Config.DB_PORT

	// DSN(Data Source Name)を作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
