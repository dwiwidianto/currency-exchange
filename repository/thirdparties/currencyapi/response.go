package currencyapi

type ResponseCurrencyApi struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	BaseCurrency string `json:"base_currency"`
}
