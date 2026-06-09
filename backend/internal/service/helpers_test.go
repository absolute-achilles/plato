package service

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/stretchr/testify/require"

	"github.com/absolute-achilles/plato/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	POSTGRES_IMAGE string = "postgres:18.4"
)

// "POSTGRES_USER":     "user",
// "POSTGRES_PASSWORD": "pass",
// "POSTGRES_DB":       "plato",
func createPosgresContainer(t *testing.T) (connectionString string, cleanup func()) {
	ctx := context.Background()

	postgresC, err := postgres.Run(ctx,
		POSTGRES_IMAGE,
		postgres.WithDatabase("plato"),
		postgres.WithUsername("user"),
		postgres.WithPassword("pass"),
		postgres.BasicWaitStrategies(),
	)
	require.NoError(t, err)

	connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	d, err := iofs.New(migrations.FS, ".")
	require.NoError(t, err)

	// golang-migrate expects the driver format (e.g., pgx5 or postgres)
	migrator, err := migrate.NewWithSourceInstance("iofs", d, connStr)
	require.NoError(t, err)

	defer migrator.Close()

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		t.Fatalf("failed to run migrations up: %s", err)
	}

	cleanup = func() {
		testcontainers.CleanupContainer(t, postgresC)
	}

	return connStr, cleanup
}
