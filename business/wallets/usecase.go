package wallets

import (
	"context"
	"currency-exchange/helpers"
	"time"
)

type WalletUseCase struct {
	repo WalletsUseCaseInterfaces
	ctx  time.Duration
}

func NewUseCase(walletRepo WalletsRepoInterfaces, contextTimeout time.Duration) WalletsUseCaseInterfaces {
	return &WalletUseCase{
		repo: walletRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *WalletUseCase) GetAllWallets(ctx context.Context) ([]Domain, error) {
	wallet, err := usecase.repo.GetAllWallets(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return wallet, nil
}

func (usecase *WalletUseCase) CreateWallets(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Total == 0 {
		return Domain{}, helpers.ErrEmailNotValid
	}
	wallet, err := usecase.repo.CreateWallets(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return wallet, nil

}

func (usecase *WalletUseCase) DeleteWallets(ctx context.Context, id int) error {
	err := usecase.repo.DeleteWallets(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
