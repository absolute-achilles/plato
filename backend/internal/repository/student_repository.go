package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jmoiron/sqlx"
)

type StudentRepository interface {
	// student struct will be filled with the resulted ID
	Create(ctx context.Context, student *domain.Student) error
}

type studentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepository{db: db}
}

// teacher struct will be filled with the resulted ID
func (r *studentRepository) Create(ctx context.Context, student *domain.Student) error {
	if student == nil {
		return fmt.Errorf("Empty Student Request")
	}

	hashedPassword, err := utils.HashPassword(student.HashPassword)
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
		VALUES ($1, $2, $3, 'STUDENT')
		RETURNING id
	`

	err = tx.QueryRowContext(ctx, insertUserQuery, student.Name, student.Email, hashedPassword).Scan(&student.ID)
	if err != nil {
		switch {
		// Duplicate
		case strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "email"):
			return fmt.Errorf("Failed to insert user: %w", ErrDuplicateEmail)
		// Duplicate
		case strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "name"):
			return fmt.Errorf("Failed to insert user: %w", ErrDuplicateName)

		// Duplicate
		case strings.Contains(err.Error(), "23505"):
			return fmt.Errorf("Failed to insert user: %w", domain.ErrDuplicate)
		default:
			return fmt.Errorf("Failed to insert user: %w", err)
		}
	}

	studentQuery := `INSERT INTO student (user_id) VALUES ($1)`
	_, err = tx.ExecContext(ctx, studentQuery, student.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert student: %w", err)
	}

	return tx.Commit()
}
