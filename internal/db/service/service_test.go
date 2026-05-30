package service

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func mustExecSQL(t *testing.T, db *sql.DB, path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("read schema %s: %v", path, err)
	}
	if _, err := db.Exec(string(b)); err != nil {
		t.Fatalf("exec schema %s: %v", path, err)
	}
}

func TestDoTransaction_DoubleEntry(t *testing.T) {
	ctx := context.Background()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		t.Fatalf("pragma: %v", err)
	}

	// apply schema files (schema is in ../schema relative to this package)
	base := filepath.Join("..", "schema")
	mustExecSQL(t, db, filepath.Join(base, "account.sql"))
	mustExecSQL(t, db, filepath.Join(base, "transactions.sql"))
	mustExecSQL(t, db, filepath.Join(base, "entries.sql"))

	svc := NewSQLService(db)

	from, err := svc.CreateAccount(ctx, "From")
	if err != nil {
		t.Fatalf("create from: %v", err)
	}
	to, err := svc.CreateAccount(ctx, "To")
	if err != nil {
		t.Fatalf("create to: %v", err)
	}

	amount := 123.45
	tx, err := svc.DoTransaction(ctx, amount, "test transfer", from.ID, to.ID)
	if err != nil {
		t.Fatalf("do transaction: %v", err)
	}
	_ = tx

	// verify entries
	entries, err := svc.q.ListEntries(ctx)
	if err != nil {
		t.Fatalf("list entries: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	var sum float64
	for _, e := range entries {
		sum += e.Amount
	}
	if sum != 0 {
		t.Fatalf("entries do not balance, sum=%v", sum)
	}
}
