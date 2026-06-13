package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ParentRepository interface {
	// parent struct will be filled with the resulted ID
	Create(ctx context.Context, parent *domain.Parent) error
}

type parentRepository struct {
	db *pgxpool.Pool
}

func NewParentRepository(db *pgxpool.Pool) ParentRepository {
	return &parentRepository{db: db}
}

// teacher struct will be filled with the resulted ID
func (r *parentRepository) Create(ctx context.Context, parent *domain.Parent) error {
	if parent == nil {
		return fmt.Errorf("Empty Parent Request")
	}

	hashedPassword, err := utils.HashPassword(parent.HashPassword)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	insertUserQuery := `
		INSERT INTO users (username, email, hash_password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err = tx.QueryRow(
		ctx,
		insertUserQuery,
		parent.Username,
		parent.Email,
		hashedPassword,
		domain.RoleParent,
	).Scan(&parent.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	parentQuery := `INSERT INTO parents (user_id) VALUES ($1)`
	_, err = tx.Exec(ctx, parentQuery, parent.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert parent: %w", err)
	}

	return tx.Commit(ctx)
}
