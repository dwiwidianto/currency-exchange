package wallets

import (
	"time"

	"golang.org/x/net/context"
)

type WalletUsaecase struct {
	Repo    WalletRepoInterfaces
	Timeout time.Duration
}

func NewWalletUsecase(repo WalletRepoInterfaces, timeout time.Duration) *WalletUsaecase {
	return &WalletUsaecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (usecase *WalletUsaecase) Create(ctx context.Context, domain Domain) error {
	err := usecase.Repo.Create(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *WalletUsaecase) GetByUserId(ctx context.Context, userId int) (Domain, error) {
	wallet, err := usecase.Repo.GetByUserId(ctx, userId)
	if err != nil {
		return Domain{}, err
	}
	return wallet, nil
}

// func (usecase *WalletUsaecase) UpdateByUserId(ctx context.Context, domain Domain) (Domain, error) {
// 	data, err := usecase.Repo.UpdateByUserId(ctx, domain)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	data.TransactionNominal = domain.TransactionNominal
// 	data.TransactionId = domain.TransactionId

// 	return data, nil
// }
