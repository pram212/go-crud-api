package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    ID    uint   `gorm:"primaryKey" json:"id"`
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    Posts  []Post  `json:"posts" gorm:"foreignKey:UserID"`
}