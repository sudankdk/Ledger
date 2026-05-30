run :
	go build -o ledger cmd/ledger/main.go  
	./ledger


include .env

migrate-up:
	goose -dir internal/db/migration sqlite3 "$(DB_URL)" up

migrate-down:
	goose -dir internal/db/migration sqlite3 "$(DB_URL)" down

migration:
	goose -dir internal/db/migration create $(name) sql