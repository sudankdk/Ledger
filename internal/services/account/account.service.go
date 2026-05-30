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

func (a *AccountService) ListAccounts(ctx context.Context) ([]db.Account, error) {
	return a.db.ListAccounts(ctx)
}

func (a *AccountService) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	return a.db.GetAccount(ctx, id)
}
