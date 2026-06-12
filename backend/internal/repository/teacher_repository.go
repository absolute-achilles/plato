package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jmoiron/sqlx"
)

type TeacherRepository interface {
	// teacher struct will be filled with the resulted ID
	Create(ctx context.Context, teacher *domain.Teacher) error
}

type teacherRepository struct {
	db *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

// teacher struct will be filled with the resulted ID
func (r *teacherRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	if teacher == nil {
		return fmt.Errorf("Empty Teacher Request")
	}

	hashedPassword, err := utils.HashPassword(teacher.HashPassword)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %w", err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	insertUserQuery := `
		INSERT INTO users (username, email, hash_password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err = tx.QueryRowContext(
		ctx,
		insertUserQuery,
		teacher.Username,
		teacher.Email,
		hashedPassword,
		domain.RoleTeacher,
	).Scan(&teacher.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	teacherQuery := `INSERT INTO teachers (user_id) VALUES ($1)`
	_, err = tx.ExecContext(ctx, teacherQuery, teacher.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert teacher: %w", err)
	}

	return tx.Commit()
}
