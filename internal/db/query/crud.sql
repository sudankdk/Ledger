-- name: CreateAccount :one
INSERT INTO account (name)
VALUES (?)
RETURNING id, name, created_at;

-- name: GetAccount :one
SELECT id, name, created_at
FROM account
WHERE id = ?;

-- name: ListAccounts :many
SELECT id, name, created_at
FROM account
ORDER BY id;

-- name: UpdateAccount :one
UPDATE account
SET name = ?
WHERE id = ?
RETURNING id, name, created_at;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;

-- name: CreateTransaction :one
INSERT INTO "transaction" (description)
VALUES (?)
RETURNING id, description, created_at;

-- name: GetTransaction :one
SELECT id, description, created_at
FROM "transaction"
WHERE id = ?;

-- name: ListTransactions :many
SELECT id, description, created_at
FROM "transaction"
ORDER BY id;

-- name: UpdateTransaction :one
UPDATE "transaction"
SET description = ?
WHERE id = ?
RETURNING id, description, created_at;

-- name: DeleteTransaction :exec
DELETE FROM "transaction"
WHERE id = ?;

-- name: CreateEntry :one
INSERT INTO entries (transaction_id, account_id, amount)
VALUES (?, ?, ?)
RETURNING id, transaction_id, account_id, amount, created_at;

-- name: GetEntry :one
SELECT id, transaction_id, account_id, amount, created_at
FROM entries
WHERE id = ?;

-- name: ListEntries :many
SELECT id, transaction_id, account_id, amount, created_at
FROM entries
ORDER BY id;

-- name: UpdateEntry :one
UPDATE entries
SET transaction_id = ?, account_id = ?, amount = ?
WHERE id = ?
RETURNING id, transaction_id, account_id, amount, created_at;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = ?;