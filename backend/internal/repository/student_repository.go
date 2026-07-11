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

type StudentRepository interface {
	// student struct will be filled with the resulted ID
	Create(ctx context.Context, student *domain.Student) error
}

type studentRepository struct {
	db *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) StudentRepository {
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

	gradeLevel := student.GradeLevel
	if gradeLevel == "" {
		gradeLevel = domain.GradeLevel1
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
		student.Username,
		student.Email,
		hashedPassword,
		domain.RoleStudent,
		student.PhoneNumber,
	).Scan(&student.ID)
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

	studentQuery := `INSERT INTO students (user_id, grade_level) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, studentQuery, student.ID, gradeLevel)
	if err != nil {
		return fmt.Errorf("Failed to insert student: %w", err)
	}

	return tx.Commit(ctx)
}
