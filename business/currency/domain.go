package currency

import "context"

type Domain struct {
	Symbol       string `json:"symbol`
	Name         string `json:"name"`
	BaseCurrency string `json:"base_currency"`
}

type Usecase interface {
	GetAll(ctx context.Context) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) (Domain, error)
}
