package transactions

import (
	"context"
	"errors"
)

type TransactionUseCase struct {
	repo TransactionRepoInterfaces
}

func NewUseCase(transactionRepo TransactionRepoInterfaces) TransactionUseCaseInterfaces {
	return &TransactionUseCase{
		repo: transactionRepo,
	}
}

func (usecase *TransactionUseCase) CreateTranasaction(domain Domain) (Domain, error) {
	if domain.BaseCurrency == "" {
		return Domain{}, errors.New("KK id cannot be empty")
	}
	if domain.SwapCurrency == 0 {
		return Domain{}, errors.New("Nik cannot be empty")
	}
	if domain.Total == 0 {
		return Domain{}, errors.New("Nama cannot be empty")
	}

	penduduk, err := usecase.repo.CreateTranasaction(domain)
	if err != nil {
		return Domain{}, err
	}
	return penduduk, nil
}

func (usecase *TransactionUseCase) GetAllTransaction(ctx context.Context) ([]Domain, error) {
	data, err := usecase.repo.GetAllTransaction(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (usecase *TransactionUseCase) DeleteTranasction(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteTranasction(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
