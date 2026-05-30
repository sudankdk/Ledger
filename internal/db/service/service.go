package service

import (
	"context"
	"database/sql"

	db "github.com/sudankdk/ledger/internal/db/sqlc"
)

type SQLService struct {
	q *db.Queries
}

func NewSQLService(dbtx *sql.DB) *SQLService {
	return &SQLService{q: db.New(dbtx)}
}

// Account methods
func (s *SQLService) CreateAccount(ctx context.Context, name string) (db.Account, error) {
	return s.q.CreateAccount(ctx, name)
}
func (s *SQLService) ListAccounts(ctx context.Context) ([]db.Account, error) {
	return s.q.ListAccounts(ctx)
}
func (s *SQLService) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	return s.q.GetAccount(ctx, id)
}

// Transaction
func (s *SQLService) AddTransaction(ctx context.Context, description string) (db.Transaction, error) {
	var desc sql.NullString
	if description != "" {
		desc = sql.NullString{String: description, Valid: true}
	} else {
		desc = sql.NullString{Valid: false}
	}
	return s.q.CreateTransaction(ctx, desc)
}

// Entry
func (s *SQLService) CreateEntry(ctx context.Context, transactionID int64, accountID int64, amount float64) (db.Entry, error) {
	arg := db.CreateEntryParams{
		TransactionID: transactionID,
		AccountID:     accountID,
		Amount:        amount,
	}
	return s.q.CreateEntry(ctx, arg)
}
