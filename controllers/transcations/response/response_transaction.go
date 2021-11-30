package response

import (
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
}
