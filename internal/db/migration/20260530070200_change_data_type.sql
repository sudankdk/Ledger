-- +goose Up
-- Rename table and migrate entries.amount to REAL
PRAGMA foreign_keys=off;

ALTER TABLE "transaction" RENAME TO transactions;

-- Recreate entries table with amount as REAL
CREATE TABLE IF NOT EXISTS entries_new (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	transaction_id INTEGER NOT NULL,
	account_id INTEGER NOT NULL,
	amount REAL NOT NULL,
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (transaction_id) REFERENCES transactions (id) ON DELETE CASCADE,
	FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE
);

INSERT INTO entries_new (id, transaction_id, account_id, amount, created_at)
SELECT id, transaction_id, account_id, amount, created_at FROM entries;

DROP TABLE entries;
ALTER TABLE entries_new RENAME TO entries;

PRAGMA foreign_keys=on;

-- +goose Down
PRAGMA foreign_keys=off;

-- Rename transactions back to "transaction" first
ALTER TABLE transactions RENAME TO "transaction";

-- Recreate entries with amount as INTEGER
CREATE TABLE IF NOT EXISTS entries_old (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	transaction_id INTEGER NOT NULL,
	account_id INTEGER NOT NULL,
	amount INTEGER NOT NULL,
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (transaction_id) REFERENCES "transaction" (id) ON DELETE CASCADE,
	FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE
);

INSERT INTO entries_old (id, transaction_id, account_id, amount, created_at)
SELECT id, transaction_id, account_id, CAST(amount AS INTEGER), created_at FROM entries;

DROP TABLE entries;
ALTER TABLE entries_old RENAME TO entries;

PRAGMA foreign_keys=on;
