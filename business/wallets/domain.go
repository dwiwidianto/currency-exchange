package wallets

import "context"

type Domain struct {
	Id                 int
	UserId             int
	Total              float64
	TransactionNominal float64
	TransactionId      int
	Nominal            float64
}

type WalletUsecaseInterfaces interface {
	GetByUserId(ctx context.Context, UserId int) (Domain, error)
	// UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}
type WalletRepoInterfaces interface {
	GetByUserId(ctx context.Context, UserId int) (Domain, error)
	// UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}
