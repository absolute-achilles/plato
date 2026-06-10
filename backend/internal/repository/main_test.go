package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/absolute-achilles/plato/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	POSTGRES_IMAGE string = "postgres:18.4"
)

var TestDBConnString string

// TestMain wraps all tests in this package.
func TestMain(m *testing.M) {
	ctx := context.Background()

	postgresC, err := createPostgresContainer(ctx)
	if err != nil {
		log.Fatalf("failed to setup postgres test container: %v", err)
	}

	// Ensure the container is terminated when all tests finish.
	defer func() {
		if err := testcontainers.TerminateContainer(postgresC); err != nil {
			log.Printf("failed to terminate postgres container: %v", err)
		}
	}()

	connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("could not get DB connection string: %v", err)
	}
	TestDBConnString = connStr

	os.Exit(m.Run())
}

func createPostgresContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	postgresC, err := postgres.Run(ctx,
		POSTGRES_IMAGE,
		postgres.WithDatabase("plato"),
		postgres.WithUsername("user"),
		postgres.WithPassword("pass"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", d, connStr)
	if err != nil {
		return nil, err
	}
	defer migrator.Close()

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return postgresC, nil
}
