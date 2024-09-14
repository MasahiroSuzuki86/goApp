package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Done    bool   `json:"done" gorm:"default:false"`
}
