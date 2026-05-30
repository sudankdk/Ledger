package account

import (
	"context"

	db "github.com/sudankdk/ledger/internal/db/sqlc"
)

type AccountService struct {
	db *db.Queries
}

func NewAccountService(db *db.Queries) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, name string) (db.Account, error) {
	return a.db.CreateAccount(ctx, name)
}
