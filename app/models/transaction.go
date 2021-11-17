package models

import (
	"time"
)

type Transaction struct {
	ID               string `gorm:"size:20;not null;uniqueIndex;primary_key"`
	TotalTransaction string `gorm:"size: 255;"`
	CreatedAt        time.Time
	UpdateAt         time.Time
}
