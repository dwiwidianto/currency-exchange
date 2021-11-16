package models

import (
	"time"
)

type Wallet struct {
	ID        string `gorm: "size: 20; not null; uniqueIndex; pramery_key"`
	User      User
	UserID    string `gorm: "size: 20; index"`
	CreatedAt time.Time
}
