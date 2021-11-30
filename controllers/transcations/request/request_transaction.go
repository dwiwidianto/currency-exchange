package request

import "currency-exchange/business/transactions"

type CreateTransaction struct {
	BaseCurrency string  `json:""`
	SwapCurrency float64 `json:""`
	Total        float64 `json:""`
}

func (transaction *CreateTransaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		BaseCurrency: transaction.BaseCurrency,
		SwapCurrency: transaction.SwapCurrency,
		Total:        transaction.SwapCurrency,
	}
}
