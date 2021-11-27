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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCaseInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, register *Domain) (Domain, error)
	Login(domain Domain, ctx context.Context) (Domain, error)
}

type UserRepoInterface interface {
	GetAllUsers(ctx context.Context) ([]Domain, error)
	GetByIdUsers(ctx context.Context, id int) (Domain, error)
	GetByEmailUsers(ctx context.Context, email string) (Domain, error)
	CreateUsers(ctx context.Context, register *Domain) (Domain, error)
	Login(domain Domain, ctx context.Context) (Domain, error)
}
