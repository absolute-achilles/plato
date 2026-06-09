package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver — side-effect import, registers the driver
)

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	//
	// If d <= 0, connections are not closed due to a connection's age.
	ConnMaxLifetime time.Duration
}

// NewPostgres returns a ready-to-use *sqlx.DB or fails fast.
// It's just a plain function — no struct, no magic.
func NewPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("database.NewPostgres: failed to connect: %w", err)
	}

	// connection pool tuning
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Ping to verify the connection is actually alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database.NewPostgres: failed to ping: %w", err)
	}

	return db, nil
}
