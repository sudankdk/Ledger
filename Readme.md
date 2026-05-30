# Ledger CLI

	_      _____  _____  _____  _____  _____
 | |    |  ___||  ___||  ___||  ___||  ___|
 | |    | |_   | |_   | |_   | |_   | |_   
 | |___ |  _|  |  _|  |  _|  |  _|  |  _|  
 |_____|_|    |_|    |_|    |_|    |_|    

A lightweight double-entry accounting CLI and library built with Go, Cobra, SQLite, and sqlc.
This repository provides a small ledger engine, an interactive CLI with history and completion, and a service layer that enforces atomic double-entry transactions.

---

## Highlights

- Double-entry accounting model with a simple numeric sign convention (debit = positive, credit = negative).
- Interactive REPL with history and completion.
- Small, testable service layer (`internal/db/service`) with unit tests.

---

## Quick Start

1. Build the CLI:

```bash
go build -o ledger ./cmd/ledger
```

2. Run migrations to create the SQLite DB (`ledger.db`):

```bash
make migrate-up
# or
goose -dir internal/db/migration sqlite3 ledger.db up
```

3. Create accounts:

```bash
./ledger account create "Checking"
./ledger account create "Savings"
```

4. Transfer (double-entry) example:

```powershell
./ledger transfer 100 "Checking" "Savings" "Monthly savings"
```

5. Or run the interactive REPL (recommended):

```powershell
.\ledger
ledger> account create "Checking"
ledger> transfer 50 "Checking" "Savings" "move to savings"
ledger> account list
ledger> exit
```

---

## Core Concepts & Conventions

Double-entry accounting rules used here:

- Every financial action is a `transaction`.
- Each `transaction` has multiple `entries` (movements per account).
- Sum of all `entries` for a transaction must equal zero.

Sign convention (simple numeric convention used by the codebase):

- Debit entries store a positive `amount` value.
- Credit entries store a negative `amount` value.

The service method `DoTransaction(ctx, amount, desc, debitAccountID, creditAccountID)` creates a `transactions` row and two `entries`: a debit `+amount` and a credit `-amount` inside a single DB transaction.

See `internal/db/README.md` for the same convention inside the DB package.

---

## Use Cases

- Personal bookkeeping: record income/expenses and transfers between accounts.
- Importer: parse bank CSVs and call `DoTransaction` for each row to import balanced entries.
- Integrations: webhook or background jobs can post payments/invoices using `DoTransaction` to ensure balanced ledger writes.
- Reporting: build balance sheets, profit/loss reports, and reconciliation tools over `transactions` and `entries` tables.

---

## CLI / REPL

The CLI exposes commands under `ledger` (via Cobra). Key commands:

- `account create <name>` — create an account
- `account list` — list accounts
- `transfer <amount> <from> <to> [description]` — double-entry transfer

Start the interactive REPL by running the binary without arguments. The REPL supports line editing and history (stored in `.ledger_history`).

---

## Testing

Run unit tests for the service layer:

```bash
go test ./internal/db/service -v
```

There is a `TestDoTransaction_DoubleEntry` that validates two balancing entries are created.

Run all tests:

```bash
go test ./...
```

---

## Development Notes

- Service layer `internal/db/service` is intentionally small and enforces atomic commits when creating transactions and entries.
- SQL sources live in `internal/db/query/*.sql`. Generated code is under `internal/db/sqlc` (managed by `sqlc`).

To regenerate SQL code after editing `.sql` files:

```bash
sqlc generate
```

---

## Future Improvements

- Add richer reporting commands (balance, ledger, trial balance).
- Add multi-currency support and currency conversions.
- Improve REPL with context-aware completion and history cross-repo.
- Add CLI subcommands for reconciliation and importers.

---

## License

MIT
