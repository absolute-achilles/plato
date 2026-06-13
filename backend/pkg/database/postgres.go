package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DSN          string
	MaxConns     *int32
	MinIdleConns *int32

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	//
	// If d <= 0, connections are not closed due to a connection's age.
	ConnMaxLifetime *time.Duration
}

// NewPostgres returns a ready-to-use *sqlx.DB or fails fast.
func NewPostgres(cfg Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("database.NewPostgres: failed to connect: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("database.NewPostgres: failed to connect: %w", err)
	}

	if cfg.MaxConns != nil {
		config.MaxConns = *cfg.MaxConns
	}
	if cfg.MinIdleConns != nil {
		config.MinIdleConns = *cfg.MinIdleConns
	}

	if cfg.ConnMaxLifetime != nil {
		config.MaxConnLifetime = *cfg.ConnMaxLifetime
	}

	// Ping to verify the connection is actually alive
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database.NewPostgres: failed to ping: %w", err)
	}

	return pool, nil
}
