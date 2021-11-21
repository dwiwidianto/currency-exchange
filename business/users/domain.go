package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserUseCaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
}
