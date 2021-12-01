package response

import "currency-exchange/business/wallets"

type WalletCreateResponse struct {
	ID            int     `json:"id"`
	Total         float64 `json:"total"`
	UserID        int     `json:"user_id"`
	TransactionID int     `json:"transaction_id"`
}

func FromDomainToCreateResponse(domain wallets.Domain) WalletCreateResponse {
	return WalletCreateResponse{
		ID:            domain.ID,
		Total:         domain.Total,
		UserID:        domain.UserID,
		TransactionID: domain.TransactionID,
	}
}
