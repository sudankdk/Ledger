package transaction

import (
	"context"
	"database/sql"

	db "github.com/sudankdk/ledger/internal/db/sqlc"
)

type TransactionService struct {
	db *db.Queries
}

func NewTransactionService(db *db.Queries) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (t *TransactionService) AddTransaction(ctx context.Context, description string) (db.Transaction, error) {
	desc := sql.NullString{String: description, Valid: description != ""}
	return t.db.CreateTransaction(ctx, desc)
}
