package currencyapi

type ResponseCurrencyApi struct {
	Query struct {
		Apikey       string `json:"apikey"`
		BaseCurrency string `json:"base_currency"`
		Timestamp    int    `json:"timestamp"`
	} `json:"query"`
	Data struct {
		Usd float64 `json:"USD"`
		Jpy float64 `json:"JPY"`
		Sgd float64 `json:"SGD"`
	} `json:"data"`
}
