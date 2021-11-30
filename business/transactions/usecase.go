package transactions

import (
	"context"
	"time"
)

type TransactionUsecase struct {
	Repo    TransactionRepo
	Timeout time.Duration
}

func NewTransactionUsecase(repo TransactionRepo, timeout time.Duration) *TransactionUsecase {
	return &TransactionUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (usecase *TransactionUsecase) GetTransactionById(ctx context.Context, id int) (Domain, error) {
	d, err := usecase.Repo.GetTransactionById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return d, nil
}

func (usecase *TransactionUsecase) CreateTransaction(ctx context.Context, domain Domain) (Domain, error) {
	d, err := usecase.Repo.CreateTransaction(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return d, nil
}
