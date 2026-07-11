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

type TeacherRepository interface {
	// teacher struct will be filled with the resulted ID
	Create(ctx context.Context, teacher *domain.Teacher) error
}

type teacherRepository struct {
	db *pgxpool.Pool
}

func NewTeacherRepository(db *pgxpool.Pool) TeacherRepository {
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

	department := teacher.Department
	if department == "" {
		department = domain.DepartmentOther
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
		teacher.Username,
		teacher.Email,
		hashedPassword,
		domain.RoleTeacher,
		teacher.PhoneNumber,
	).Scan(&teacher.ID)
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

	teacherQuery := `INSERT INTO teachers (user_id, department) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, teacherQuery, teacher.ID, department)
	if err != nil {
		return fmt.Errorf("Failed to insert teacher: %w", err)
	}

	return tx.Commit(ctx)
}
