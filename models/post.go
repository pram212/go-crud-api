package models

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    UserID  uint `json:"user_id"`   // foreign key
    User    User `gorm:"foreignKey:UserID"` // relasi ke User
}