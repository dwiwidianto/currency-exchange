package currencyapi

import (
	"context"
	"currency-exchange/business/currency"
	"encoding/json"
	"net/http"
)

type Currencyapi struct {
	httpClient http.Client
}

func NewCurrencyApi() currency.Repository {
	return &Currencyapi{
		httpClient: http.Client{},
	}
}

func (repo *Currencyapi) GetAll(ctx context.Context) (res currency.Domain, err error) {
	req, _ := http.NewRequest("GET", "https://freecurrencyapi.net/api/v2/latest?apikey=68b7ee20-3f28-11ec-9fbd-d327ca8d9059&base_currency=IDR", nil)
	resp, err := repo.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	bind := ResponseCurrencyApi{}
	err := json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		return res, err
	}
	res = currency.Domain{}
	for _, datafrom := range bind {
		res.Query = datafrom.Query
		res.Data = datafrom.Data
	}
	return res, nil
}
