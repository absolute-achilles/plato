package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jmoiron/sqlx"
)

type AdminRepository interface {
	// admin struct will be filled with the resulted ID
	Create(ctx context.Context, admin *domain.Admin) error
}

type adminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) AdminRepository {
	return &adminRepository{db: db}
}

// teacher struct will be filled with the resulted ID
func (r *adminRepository) Create(ctx context.Context, admin *domain.Admin) error {
	if admin == nil {
		return fmt.Errorf("Empty Admin Request")
	}

	hashedPassword, err := utils.HashPassword(admin.HashPassword)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	insertUserQuery := `
		INSERT INTO "user" (name, email, hash_password, role)
		VALUES ($1, $2, $3, 'ADMIN')
		RETURNING id
	`

	err = tx.QueryRowContext(ctx, insertUserQuery, admin.Name, admin.Email, hashedPassword).Scan(&admin.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	adminQuery := `INSERT INTO admin (user_id) VALUES ($1)`
	_, err = tx.ExecContext(ctx, adminQuery, admin.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert admin: %w", err)
	}

	return tx.Commit()
}
