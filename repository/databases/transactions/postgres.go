package transactions

import (
	"context"
	"currency-exchange/business/transactions"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(gormDb *gorm.DB) transactions.TransactionRepoInterfaces {
	return &TransactionRepository{
		db: gormDb,
	}
}

func (repo *TransactionRepository) InsertNewDataPenduduks(domain transactions.Domain) (transactions.Domain, error) {
	userInput := FromDomain(domain)

	if err := repo.db.Create(&userInput).Error; err != nil {
		return transactions.Domain{}, err
	}

	return userInput.ToDomain(), nil
}

func (repo *TransactionRepository) GetAllTransaction(ctx context.Context) ([]transactions.Domain, error) {
	var transactionsDB []Transactions

	result := repo.db.Preload("User").Find(&transactionsDB)

	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}

	return ListTransactionToDomain(transactionsDB), nil
}

func (repo *TransactionRepository) DeleteDataPenduduks(ctx context.Context, id uint) error {
	data := Transactions{}
	err := repo.db.Delete(&data, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("no id")
	}
	return nil
}
