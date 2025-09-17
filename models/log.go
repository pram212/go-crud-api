package models

import "time"

type Log struct {
    ID        uint      `gorm:"primaryKey"`
    Action    string
    Detail    string
    CreatedAt time.Time
}
