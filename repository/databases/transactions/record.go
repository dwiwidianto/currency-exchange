package transactions

import (
	"currency-exchange/business/transactions"
	"currency-exchange/repository/databases/users"
	"time"
)

type Transactions struct {
	Id                int `gorm:"primaryKey"`
	UserId            int
	User              users.User `gorm:"foreignKey:user_id"`
	BaseCurrency      string
	Currency          string
	DatetimeRequest   time.Time
	DatetimeVerify    *time.Time
	DatetimeCompleted *time.Timer
}

func TransactionsFromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		UserId:          domain.UserId,
		BaseCurrency:    domain.BaseCurrency,
		Currency:        domain.Currency,
		DatetimeRequest: time.Now().Local(),
	}
}

// func (rec *Transactions) TransactionsToDomain() (res *transactions.Domain) {
// 	if rec != nil {
// 		res = &transactions.Domain{
// 			UserId:          rec.UserId,
// 			BaseCurrency:    rec.BaseCurrency,
// 			Currency:        rec.Currency,
// 			DatetimeRequest: time.Now().Local(),
// 		}
// 	}
// 	return res
// }

func (rec *Transactions) TransactionsToDomain() transactions.Domain {
	return transactions.Domain{
		Id:              rec.Id,
		UserId:          rec.UserId,
		BaseCurrency:    rec.BaseCurrency,
		Currency:        rec.Currency,
		DatetimeRequest: time.Now().Local(),
	}
}
