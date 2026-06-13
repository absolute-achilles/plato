package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository interface {
	// admin struct will be filled with the resulted ID
	Create(ctx context.Context, admin *domain.Admin) error
}

type adminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) AdminRepository {
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
		admin.Username,
		admin.Email,
		hashedPassword,
		domain.RoleAdmin,
	).Scan(&admin.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	adminQuery := `INSERT INTO admins (user_id) VALUES ($1)`
	_, err = tx.Exec(ctx, adminQuery, admin.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert admin: %w", err)
	}

	return tx.Commit(ctx)
}
