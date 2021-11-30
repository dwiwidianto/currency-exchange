package wallets

import (
	"currency-exchange/business/users"
	"currency-exchange/business/wallets"
	"time"
)

type Wallets struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	// Transaction transactions.Transactions `gorm:"foreignKey:transaction_id"`
	Total    float64
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}

func (data *Wallets) WalletsToWalletsDomain() wallets.Domain {
	return wallets.Domain{
		Id:     data.Id,
		UserId: data.UserId,
		Total:  data.Total,
	}
}

func (data *Wallets) WalletToUsersWalletDomain() users.Wallets {
	return users.Wallets{
		Id:    data.Id,
		Total: data.Total,
	}
}

func WalletsFromDomain(domain wallets.Domain) Wallets {
	return Wallets{
		Id:     domain.Id,
		UserId: domain.UserId,
		Total:  domain.Total,
	}
}

func WalletsToListDomain(data []Wallets) []wallets.Domain {
	var list []wallets.Domain
	for _, v := range data {
		list = append(list, v.WalletsToWalletsDomain())
	}
	return list
}
