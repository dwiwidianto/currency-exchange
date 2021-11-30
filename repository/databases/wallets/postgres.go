package wallets

import (
	"context"
	"currency-exchange/business/wallets"
	"errors"

	"gorm.io/gorm"
)

type WalletsRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(gormDb *gorm.DB) wallets.WalletRepoInterfaces {
	return &WalletsRepository{
		db: gormDb,
	}
}

func (repo *WalletsRepository) GetByUserId(ctx context.Context, userId int) (wallets.Domain, error) {
	var data Wallets
	err := repo.db.Find(&data, "user_id=?", userId)
	if err != nil {
		return wallets.Domain{}, err.Error
	}
	if err.RowsAffected == 0 {
		return wallets.Domain{}, errors.New("user id not found")
	}
	return data.WalletsToWalletsDomain(), nil
}

func (repo *WalletsRepository) Create(ctx context.Context, domain wallets.Domain) error {
	data := WalletsFromDomain(domain)
	err := repo.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
