package entries

import (
	"context"

	db "github.com/sudankdk/ledger/internal/db/sqlc"
)

type EntryService struct {
	db *db.Queries
}

func NewEntryService(db *db.Queries) *EntryService {
	return &EntryService{
		db: db,
	}
}

func (e *EntryService) CreateEntry(ctx context.Context, transactionID int64, accountID int64, amount float64) (db.Entry, error) {
	return e.db.CreateEntry(ctx, db.CreateEntryParams{
		TransactionID: transactionID,
		AccountID:     accountID,
		Amount:        amount,
	})
}
