package transactions

import (
	"context"
	"currency-exchange/business/transactions"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewPostgresTransactionRepo(gormDb *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		db: gormDb,
	}
}

func (repo *TransactionRepo) GetTransactionById(ctx context.Context, id int) (transactions.Domain, error) {
	model := Transactions{}
	err := repo.db.First(&model, "id=?", id)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return model.TransactionsToDomain(), nil
}

func (repo *TransactionRepo) CreateTransaction(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	data := TransactionsFromDomain(domain)
	err := repo.db.Create(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}
