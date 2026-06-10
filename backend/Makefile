# ==============================================================================
# Database Variables
# Replace these with your actual database credentials
# ==============================================================================

DB_USER := user
DB_PASSWORD := pass
DB_HOST := localhost
DB_PORT := 5432
DB_NAME := plato
DB_URL := "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATION_PATH := db/migrations

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==============================================================================
# Migration Commands
# ==============================================================================

## run: run golang program
.PHONY: run
run:
	go run cmd/api/main.go

## test: run all test golang program verbose (no cache)
.PHONY: test
test:
	go test ./... -v --count=1

# ==============================================================================
# Migration Commands
# ==============================================================================

## migrate-up: apply all up migrations
.PHONY: migrate-up
migrate-up:
	migrate -path $(MIGRATION_PATH) -database $(DB_URL) up

## migrate-up-1: apply just the next up migration
.PHONY: migrate-up-1
migrate-up-1:
	migrate -path $(MIGRATION_PATH) -database $(DB_URL) up 1

## migrate-down: apply all down migrations (careful: drops everything)
.PHONY: migrate-down
migrate-down:
	migrate -path $(MIGRATION_PATH) -database $(DB_URL) down

## migrate-down-1: apply just the previous down migration (rollback one step)
.PHONY: migrate-down-1
migrate-down-1:
	migrate -path $(MIGRATION_PATH) -database $(DB_URL) down 1

## migrate-create name=$1: create a new migration pair with timestamp
.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Use: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)
