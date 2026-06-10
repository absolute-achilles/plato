package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jmoiron/sqlx"
)

type ParentRepository interface {
	// student struct will be filled with the resulted ID
	Create(ctx context.Context, parent *domain.Parent) error
}

type parentRepository struct {
	db *sqlx.DB
}

func NewParentRepository(db *sqlx.DB) ParentRepository {
	return &parentRepository{db: db}
}

// teacher struct will be filled with the resulted ID
func (r *parentRepository) Create(ctx context.Context, parent *domain.Parent) error {
	if parent == nil {
		return fmt.Errorf("Empty Parent Request")
	}

	hashedPassword, err := utils.HashPassword(parent.Password)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	insertUserQuery := `
		INSERT INTO "user" (name, email, password, role)
		VALUES ($1, $2, $3, 'PARENT')
		RETURNING id
	`

	err = tx.QueryRowContext(ctx, insertUserQuery, parent.Name, parent.Email, hashedPassword).Scan(&parent.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	parentQuery := `INSERT INTO parent (user_id) VALUES ($1)`
	_, err = tx.ExecContext(ctx, parentQuery, parent.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert parent: %w", err)
	}

	return tx.Commit()
}
