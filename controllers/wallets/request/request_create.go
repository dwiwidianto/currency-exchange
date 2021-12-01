package request

import "currency-exchange/business/wallets"

type CreateWallet struct {
	Total         float64 `json:"total"`
	UserID        int     `json:"user_id"`
	TransactionID int     `json:"transaction_id"`
}

func (cw *CreateWallet) ToDomain() *wallets.Domain {
	return &wallets.Domain{
		Total:         cw.Total,
		UserID:        cw.UserID,
		TransactionID: cw.TransactionID,
	}
}
