package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID               string `gorm:"size:20; not null;uniqueIndex;primary_key"`
	User             User
	UserID           string `gorm:"size: 20; index"`
	TotalTransaction string `gorm:"size: 255;"`
	CreatedAt        time.Time
	UpdateAt         time.Time
	DeletedAt        gorm.DeletedAt
}
