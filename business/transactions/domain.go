package transactions

import (
	"context"
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
}

type TransactionUseCaseInterfaces interface {
	CreateTranasaction(domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	DeleteTranasction(ctx context.Context, id uint) error
}

type TransactionRepoInterfaces interface {
	CreateTranasaction(domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
	DeleteTranasction(ctx context.Context, id uint) error
}
