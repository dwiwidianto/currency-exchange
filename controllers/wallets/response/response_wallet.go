package response

import (
	"currency-exchange/business/wallets"
	"time"

	"gorm.io/gorm"
)

type WalletResponse struct {
	ID        int            `json:"id"`
	Total     float64        `json:"total"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromWalletDomain(domain wallets.Domain) WalletResponse {
	return WalletResponse{
		ID:        domain.ID,
		Total:     domain.Total,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromWalletListDomain(domain []wallets.Domain) []WalletResponse {
	var list []WalletResponse
	for _, v := range domain {
		list = append(list, FromWalletDomain(v))
	}
	return list
}
