package transactions

import (
	"context"
	"currency-exchange/repository/databases/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int
	BaseCurrency string
	SwapCurrency float64
	Total        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	User         users.User
	UserID       int
}

type TransactionUseCaseInterfaces interface {
	CreateTransaction(domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	DeleteTransaction(ctx context.Context, id uint) error
}

type TransactionRepoInterfaces interface {
	CreateTransaction(domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	DeleteTransaction(ctx context.Context, id uint) error
}
