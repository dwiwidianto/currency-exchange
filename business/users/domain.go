package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCaseInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, register *Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, string, error)
	Delete(ctx context.Context, id int) error
}

type UserRepoInterface interface {
	GetAllUsers(ctx context.Context) ([]Domain, error)
	GetByIdUsers(ctx context.Context, id int) (Domain, error)
	GetByEmailUsers(ctx context.Context, email string) (Domain, error)
	CreateUsers(ctx context.Context, register *Domain) (Domain, error)
	DeleteUsers(ctx context.Context, id int) error
}
