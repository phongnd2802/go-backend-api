GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:12345@tcp(127.0.0.1:3301)/eCommerce"
GOOSE_MIGRATION_DIR ?= sql/schemas

up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlc:
	sqlc generate

run-dev:
	go run ./cmd/server/

run:
	docker-compose up

stop:
	docker-compose down

.PHONY: run stop up down sqlc
