package transactions

import (
	"currency-exchange/business/transactions"
	"currency-exchange/repository/databases/users"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	ID           int `gorm:"primaryKey"`
	BaseCurrency string
	SwapCurrency float64
	Total        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       int
	User         users.User `gorm:"foreignKey:UserID;references:ID"`
}

func (transaction Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:           transaction.ID,
		BaseCurrency: transaction.BaseCurrency,
		SwapCurrency: transaction.SwapCurrency,
		Total:        transaction.Total,
		CreatedAt:    transaction.UpdatedAt,
		UpdatedAt:    transaction.UpdatedAt,
		DeletedAt:    transaction.DeletedAt,
		User:         transaction.User,
		UserID:       transaction.UserID,
	}
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		ID:           domain.ID,
		BaseCurrency: domain.BaseCurrency,
		SwapCurrency: domain.SwapCurrency,
		Total:        domain.Total,
		CreatedAt:    domain.UpdatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		UserID:       domain.UserID,
	}
}

func ListTransactionToDomain(data []Transactions) []transactions.Domain {
	var list []transactions.Domain
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
