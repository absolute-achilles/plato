# Plato

## Project Shape

- Backend-only Go API. Frontend (NextJS) is planned but not present in the repo yet.
- Go module: `github.com/absolute-achilles/plato` (Go 1.26.3).
- Entrypoint: `backend/cmd/api/main.go`.
- Layered architecture: Gin handlers → services → repositories → Postgres (`pgxpool`).
  - `internal/handler/` — HTTP routes (Gin)
  - `internal/service/` — business logic
  - `internal/repository/` — DB access
  - `internal/domain/` — core models
  - `internal/dto/` — request/response structs
  - `internal/storage/` — object-store interface (S3-like)
  - `pkg/` — shared packages (database, logger, response helpers)

## Running & Building

- **Local dev server**: `cd backend && go run cmd/api/main.go` (listens on `:8080`).
- **Docker**: `docker compose up -d --build` from repo root. Compose brings up `postgres:16-alpine` + backend. Backend depends on DB healthcheck and connects via `DB_URL=postgres://user:pass@db:5432/plato?sslmode=disable`.
- **Binary build** (Dockerfile pattern): `cd backend && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o binary ./cmd/api`.

## Testing

- **Run all tests**: `cd backend && go test ./... -v --count=1` (or `make test` / `just test`).
- **Docker is required**: repository and service tests use `testcontainers-go/modules/postgres` to spin up real Postgres containers. Tests will fail if the Docker daemon is not running.
- **Test migration image**: `repository/helper_test.go` and `service/helper_test.go` both use `postgres:18.4`. Docker Compose uses `postgres:16-alpine`.
- **In-test migrations**: Tests automatically apply migrations from the embedded `migrations.FS` (`migrations/migrations.go` embeds `*.sql`).

## Migrations (Critical Gotcha)

- Migration SQL files live in `backend/migrations/`.
- Both the `Makefile` and `Justfile` set `MIGRATION_PATH := db/migrations`, which **does not exist**. Running `make migrate-up` or `just migrate-up` will fail unless you override the path to `migrations`.
- The `golang-migrate/migrate` CLI is the expected tool for manual migration commands.

## Codegen

- **Mockery**: `.mockery.yml` generates `mocks_test.go` in the same directory as the interface. Only `internal/storage` is currently configured (`all: true`). Re-run `mockery` after changing that interface.

## Notable Quirks

- `main.go` currently has the database connection block commented out. The app starts without a DB connection and only exposes a health-check endpoint.
- Logger writes to both `logs/app.log` and stdout; the `logs/` directory is created at runtime.
- No CI workflows, linting config (e.g., `.golangci.yml`), or pre-commit hooks are present yet.

## Memory (agentmemory)

Use `memory_save` and `memory_recall` to persist and retrieve project context across sessions.

- **Always save** architectural decisions, discovered gotchas, and workflow patterns.
- Use `project: "plato"` for all memories in this project.
- Use `type` field: `architecture`, `bug`, `workflow`, `pattern`, `fact`.
- Add `concepts` and `files` for precise recall later.
- **Recall at session start** to check for existing context before investigating.
