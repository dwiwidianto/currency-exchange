package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID                string `gorm:"size:20;not null;uniqueIndex;pramery_key"`
	User              User
	UserID            string `gorm:"size:20;index"`
	transaction_total string `gorm:"size:100"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
}
