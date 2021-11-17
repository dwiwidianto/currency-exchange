package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        string    `gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	UserFK    User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
