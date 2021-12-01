package wallets

import (
	"currency-exchange/business/wallets"
	"currency-exchange/repository/databases/transactions"
	"currency-exchange/repository/databases/users"
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID            int `gorm:"primaryKey"`
	Total         float64
	UserID        int
	TransactionID int
	User          users.User                `gorm:"foreignKey:UserID;references:ID"`
	Transactions  transactions.Transactions `gorm:"foreignKey:TransactionID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (rec *Wallet) ToDomain() (res *wallets.Domain) {
	if rec != nil {
		res = &wallets.Domain{
			ID:            rec.ID,
			Total:         rec.Total,
			TransactionID: rec.TransactionID,
			UserID:        rec.UserID,
			Transactions:  rec.Transactions,
			CreatedAt:     rec.CreatedAt,
			UpdatedAt:     rec.UpdatedAt,
		}
	}
	return res
}

func FromDomain(domain wallets.Domain) Wallet {
	return Wallet{
		ID:            domain.ID,
		Total:         domain.Total,
		TransactionID: domain.TransactionID,
		UserID:        domain.UserID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func ListWalletToDomain(data []Wallet) []wallets.Domain {
	var list []wallets.Domain
	for _, v := range data {
		list = append(list, *v.ToDomain())
	}
	return list
}
