package response

import (
	"currency-exchange/business/transactions"
	"time"

	"gorm.io/gorm"
)

type TransactionResponse struct {
	ID           int            `json:"id"`
	BaseCurrency string         `json:"base_currency"`
	SwapCurrency float64        `json:"swap_currnecy"`
	Total        float64        `json:"total"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
	User         interface{}    `json:"user,omitempty"`
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		ID:           domain.ID,
		BaseCurrency: domain.BaseCurrency,
		SwapCurrency: domain.SwapCurrency,
		Total:        domain.Total,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		User:         domain.User,
	}
}

func FromListDomain(domain []transactions.Domain) []TransactionResponse {
	var result []TransactionResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
