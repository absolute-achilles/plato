package repository

import (
	"context"
	"fmt"
	"strings"

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
		INSERT INTO users (username, email, hash_password, role, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err = tx.QueryRow(
		ctx,
		insertUserQuery,
		parent.Username,
		parent.Email,
		hashedPassword,
		domain.RoleParent,
		parent.PhoneNumber,
	).Scan(&parent.ID)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "email"):
			return fmt.Errorf("Failed to insert user: %w", ErrDuplicateEmail)
		case strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "username"):
			return fmt.Errorf("Failed to insert user: %w", ErrDuplicateName)
		case strings.Contains(err.Error(), "23505"):
			return fmt.Errorf("Failed to insert user: %w", domain.ErrDuplicate)
		default:
			return fmt.Errorf("Failed to insert user: %w", err)
		}
	}

	parentQuery := `INSERT INTO parents (user_id, type) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, parentQuery, parent.ID, parent.Type)
	if err != nil {
		return fmt.Errorf("Failed to insert parent: %w", err)
	}

	for _, studentID := range parent.StudentIDs {
		linkQuery := `
			INSERT INTO parent_student_links (parent_id, student_id)
			VALUES ($1, $2)
			ON CONFLICT (parent_id, student_id) DO NOTHING
		`
		if _, err := tx.Exec(ctx, linkQuery, parent.ID, studentID); err != nil {
			return fmt.Errorf("Failed to link parent to student %s: %w", studentID, err)
		}
	}

	return tx.Commit(ctx)
}
