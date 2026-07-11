set dotenv-load

backend_dir := "backend"
frontend_dir := "frontend"
compose_file := "docker-compose.yaml"

database_url := "postgres://user:pass@localhost:5432/plato?sslmode=disable"
migration_path := "migrations"

# Default: show available recipes
help:
    @just --list --unsorted

# -----------------------------------------------------------------------------
# Docker Compose (backend + Postgres)
# -----------------------------------------------------------------------------

# Start the backend and database in Docker
up:
    docker compose -f {{compose_file}} up -d --build

# Stop and remove Docker containers and volumes
down:
    docker compose -f {{compose_file}} down -v

# View backend logs
logs:
    docker compose -f {{compose_file}} logs -f backend

# -----------------------------------------------------------------------------
# Local development
# -----------------------------------------------------------------------------

# Run the backend dev server (requires a local Postgres on :5432)
backend-dev:
    cd {{backend_dir}} && DATABASE_URL={{database_url}} JWT_SECRET=dev-secret go run cmd/api/main.go

# Run the frontend dev server
frontend-dev:
    cd {{frontend_dir}} && pnpm dev

# Run backend and frontend dev servers together
dev:
    #!/usr/bin/env bash
    set -euo pipefail
    cd {{backend_dir}} && DATABASE_URL={{database_url}} JWT_SECRET=dev-secret go run cmd/api/main.go &
    BACKEND_PID=$!
    cd {{frontend_dir}} && pnpm dev &
    FRONTEND_PID=$!
    trap 'kill $BACKEND_PID $FRONTEND_PID 2>/dev/null || true' EXIT
    wait

# -----------------------------------------------------------------------------
# Testing
# -----------------------------------------------------------------------------

# Run backend service and handler tests (no Docker required)
backend-test:
    cd {{backend_dir}} && go test ./internal/service/... ./internal/handler/... -v --count=1

# Run all backend tests including repository tests (Docker required)
backend-test-all:
    cd {{backend_dir}} && go test ./... -v --count=1

# Run frontend type checks and lint
frontend-test:
    cd {{frontend_dir}} && pnpm typecheck && pnpm lint

# Run all project tests that don't require Docker
test:
    just backend-test
    just frontend-test

# -----------------------------------------------------------------------------
# Build
# -----------------------------------------------------------------------------

# Build the backend binary
backend-build:
    cd {{backend_dir}} && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o binary ./cmd/api

# Build the frontend
frontend-build:
    cd {{frontend_dir}} && pnpm build

# -----------------------------------------------------------------------------
# Lint and format
# -----------------------------------------------------------------------------

# Run Go vet and frontend lint
lint:
    cd {{backend_dir}} && go vet ./...
    cd {{frontend_dir}} && pnpm lint

# Format Go and frontend code
format:
    cd {{backend_dir}} && gofmt -w .
    cd {{frontend_dir}} && pnpm format

# -----------------------------------------------------------------------------
# Migrations (requires golang-migrate/migrate CLI)
# -----------------------------------------------------------------------------

# Apply all pending migrations
migrate-up:
    cd {{backend_dir}} && migrate -database {{database_url}} -path {{migration_path}} up

# Roll back the last migration
migrate-down:
    cd {{backend_dir}} && migrate -database {{database_url}} -path {{migration_path}} down 1

# Create a new migration pair (usage: just migrate-create add_user_profile)
migrate-create name:
    cd {{backend_dir}} && migrate create -ext sql -dir {{migration_path}} -seq {{name}}

# -----------------------------------------------------------------------------
# Short aliases
# -----------------------------------------------------------------------------

alias b := backend-test
alias f := frontend-test
alias t := test
alias l := lint
