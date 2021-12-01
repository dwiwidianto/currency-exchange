package wallets

import (
	"context"
	"currency-exchange/business/wallets"
	"errors"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewPostgresWalletRepository(gormDb *gorm.DB) wallets.WalletsRepoInterfaces {
	return &WalletRepository{
		db: gormDb,
	}
}

func (repo *WalletRepository) GetAllWallets(ctx context.Context) ([]wallets.Domain, error) {
	var data []Wallet
	err := repo.db.Find(&data)
	if err.Error != nil {
		return []wallets.Domain{}, err.Error
	}
	return ListWalletToDomain(data), nil
}

func (repo *WalletRepository) CreateWallets(ctx context.Context, createTableWallet *wallets.Domain) (wallets.Domain, error) {
	wallet := Wallet{
		Total:         createTableWallet.Total,
		UserID:        createTableWallet.UserID,
		TransactionID: createTableWallet.UserID,
	}
	err := repo.db.Create(&wallet)
	if err.Error != nil {
		return wallets.Domain{}, err.Error
	}
	return *wallet.ToDomain(), nil
}

func (repo *WalletRepository) DeleteWallets(ctx context.Context, id int) error {
	wallet := Wallet{}
	err := repo.db.Delete(&wallet, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
