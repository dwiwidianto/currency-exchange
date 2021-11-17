package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID            string `gorm:"size:20; not null;uniqueIndex;primary_key"`
	User          User
	UserID        string `gorm:"size: 20; index"`
	Transaction   Transaction
	TransactionID string `grom:"size: 255; index"`
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeletedAt     gorm.DeletedAt
}
