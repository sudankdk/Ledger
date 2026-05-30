package service

import (
	"context"
	"database/sql"

	db "github.com/sudankdk/ledger/internal/db/sqlc"
)

type SQLService struct {
	q  *db.Queries
	db *sql.DB
}

func NewSQLService(dbtx *sql.DB) *SQLService {
	return &SQLService{q: db.New(dbtx), db: dbtx}
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

func (s *SQLService) GetAccountByName(ctx context.Context, name string) (db.Account, error) {
	return s.q.GetAccountByName(ctx, name)
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

// DoTransaction performs a double-entry transaction: it creates a transaction record
// and two entries (debit and credit) with equal magnitude and opposite signs.
func (s *SQLService) DoTransaction(ctx context.Context, amount float64, desc string, debitAccountID int64, creditAccountID int64) (db.Transaction, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return db.Transaction{}, err
	}

	qtx := db.New(tx)

	var descNull sql.NullString
	if desc != "" {
		descNull = sql.NullString{String: desc, Valid: true}
	} else {
		descNull = sql.NullString{Valid: false}
	}

	transaction, err := qtx.CreateTransaction(ctx, descNull)
	if err != nil {
		_ = tx.Rollback()
		return db.Transaction{}, err
	}

	debitArg := db.CreateEntryParams{
		TransactionID: transaction.ID,
		AccountID:     debitAccountID,
		Amount:        amount,
	}
	if _, err := qtx.CreateEntry(ctx, debitArg); err != nil {
		_ = tx.Rollback()
		return db.Transaction{}, err
	}

	creditArg := db.CreateEntryParams{
		TransactionID: transaction.ID,
		AccountID:     creditAccountID,
		Amount:        -amount,
	}
	if _, err := qtx.CreateEntry(ctx, creditArg); err != nil {
		_ = tx.Rollback()
		return db.Transaction{}, err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return db.Transaction{}, err
	}

	return transaction, nil
}
