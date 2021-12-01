package request

import "currency-exchange/business/transactions"

type CreateTransaction struct {
	BaseCurrency string  `json:"base_currency"`
	SwapCurrency float64 `json:"swap_currency"`
	Total        float64 `json:"total"`
}

func (transaction *CreateTransaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		BaseCurrency: transaction.BaseCurrency,
		SwapCurrency: transaction.SwapCurrency,
		Total:        transaction.SwapCurrency,
	}
}
