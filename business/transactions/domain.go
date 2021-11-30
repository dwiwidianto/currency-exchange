package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id                int
	UserId            int
	BaseCurrency      string
	Currency          string
	DatetimeRequest   time.Time
	DatetimeVerify    *time.Time
	DatetimeCompleted *time.Timer
}
type TransactionUseCase interface {
	GetTransactionById(ctx context.Context, id int) (Domain, error)
	CreateTransaction(ctx context.Context, domain Domain) (Domain, error)
}

type TransactionRepo interface {
	GetTransactionById(ctx context.Context, id int) (Domain, error)
	CreateTransaction(ctx context.Context, domain Domain) (Domain, error)
}
