package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"size:20; not null;uniqueIndex;primary_key"`
	Email     string `gorm:"size:100; not null"`
	Name      string `gorm:"size:100; not null; uniqueIndex"`
	Password  string `gorm:"size:255; not null"`
	Phone     int    `gorm:"size 255; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
