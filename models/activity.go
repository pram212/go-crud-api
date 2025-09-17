package models

import (
	"time"
	"go-crud-api/database"
)

type Activity struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint
	Action    string
	Timestamp time.Time `gorm:"autoCreateTime"`
}

func CreateActivity(userID uint, action string) {
	activity := Activity{
		UserID: userID,
		Action: action,
	}

	database.DB.Create(&activity)
}
