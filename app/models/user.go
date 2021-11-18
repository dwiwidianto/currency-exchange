package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Phone     int       `json:"phone"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
