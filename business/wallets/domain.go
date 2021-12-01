package wallets

import (
	"context"
	"currency-exchange/repository/databases/transactions"
	"currency-exchange/repository/databases/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            int
	Total         float64
	UserID        int
	TransactionID int
	User          users.User
	Transactions  transactions.Transactions
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type WalletsUseCaseInterfaces interface {
	GetAllWallets(ctx context.Context) ([]Domain, error)
	CreateWallets(ctx context.Context, createTableWallet *Domain) (Domain, error)
	DeleteWallets(ctx context.Context, id int) error
}

type WalletsRepoInterfaces interface {
	GetAllWallets(ctx context.Context) ([]Domain, error)
	CreateWallets(ctx context.Context, createTableWallet *Domain) (Domain, error)
	DeleteWallets(ctx context.Context, id int) error
}
