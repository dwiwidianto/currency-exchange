package currency

import (
	"context"
	"time"
)

type currencyApiUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewCurrencyApiUsecase(timeout time.Duration, wr Repository) Usecase {
	return &currencyApiUsecase{
		contextTimeout: timeout,
		repo:           wr,
	}
}

func (uc *currencyApiUsecase) GetAll(ctx context.Context) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.repo.GetAll(ctx)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
