# Ledger CLI

A lightweight **double-entry accounting CLI** built with Go, Cobra, SQLite, and sqlc.
It is designed to model real accounting principles with strict transaction balancing.

---

## Core Concept

This system implements **double-entry accounting**, where:

* Every financial action is a **transaction**
* Each transaction contains multiple **entries**
* All entries must balance to zero

```
Assets = Liabilities + Equity
```

---

## Features

* Account management (create accounts)
* Double-entry transaction system
* Strict balance validation
* SQLite local storage
* Type-safe SQL with sqlc
* CLI interface powered by Cobra
* Clean layered architecture (cmd → service → repo → db)

---

## Tech Stack

* Go
* Cobra (CLI framework)
* SQLite
* sqlc (type-safe SQL generation)
* pgx / database/sql driver (SQLite driver depending on setup)

---

## Architecture

```
cmd/        → CLI layer (Cobra)
service/    → Business logic (double-entry rules)
repository/ → Database abstraction
internal/db → sqlc generated code
```

---

## Database Model

### accounts

Stores user-defined financial accounts.

### transactions

Represents a financial event.

### entries

Represents debit/credit movements linked to accounts.

---

## Workflow

### 1. Create accounts

```
ledger account create Cash
ledger account create Salary
```

### 2. Create transaction

```
ledger tx create \
  --desc "Salary May" \
  --credit Cash=5000 \
  --debit Salary=5000
```

### 3. View balances

```
ledger account balance Cash
```

---

## Rules

* Every transaction must balance:

  ```
  SUM(entries) = 0
  ```
* Credits increase value, debits decrease value
* Entries are immutable once committed (recommended design)

---

## Example

### Transaction: Salary Received

| Account | Amount |
| ------- | ------ |
| Cash    | +5000  |
| Income  | -5000  |

Result: Balanced transaction

---

## Setup

### 1. Install dependencies

```
go mod tidy
```

### 2. Install Cobra CLI

```
go install github.com/spf13/cobra-cli@latest
```

### 3. Install sqlc

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### 4. Generate SQL code

```
sqlc generate
```

### 5. Run CLI

```
go run main.go --help
```

---

## Build

```
go build -o ledger cmd/ledger/main.go
```

---

## Design Philosophy

This project follows real-world accounting principles:

* No direct account mutation
* All changes go through transactions
* Strong consistency enforced in service layer
* Database is treated as an immutable ledger store

---

## Future Improvements

* Balance sheet reports
* Income vs expense reports
* Interactive CLI mode
* Multi-currency support
* Audit logs
* Transaction rollback safety

---

## License

MIT
