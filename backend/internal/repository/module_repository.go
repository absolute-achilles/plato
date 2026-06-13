package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ModuleRepository interface {
	Get(ctx context.Context, moduleID string) (*domain.Module, error)
	Create(ctx context.Context, module *domain.Module) (id string, err error)
	Delete(ctx context.Context, moduleID string) error

	GetCourseModules(ctx context.Context, courseID string) ([]*domain.Module, error)
}

type moduleRepository struct {
	db *pgxpool.Pool
}

func NewModuleRepository(db *pgxpool.Pool) ModuleRepository {
	return &moduleRepository{db: db}
}

func (r *moduleRepository) Get(ctx context.Context, moduleID string) (*domain.Module, error) {
	query := `
	SELECT id, course_id, name, position, is_published, unlock_date, created_at,updated_at 
	FROM modules WHERE id = $1`

	rows, err := r.db.Query(ctx, query, moduleID)
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.Get: %w", err)
	}
	module, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Module])
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.Get: %w", err)
	}

	return module, nil
}

func (r *moduleRepository) Create(ctx context.Context, module *domain.Module) (id string, err error) {
	if module == nil {
		return "", fmt.Errorf("moduleRepository.Create: Empty Module Request")
	}

	if module.CourseID == "" {
		return "", fmt.Errorf("moduleRepository.Create: Empty Course ID")
	}

	insertQuery := `
		INSERT INTO modules (course_id, name, position, is_published, unlock_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	if err := r.db.QueryRow(
		ctx,
		insertQuery,
		module.CourseID,
		module.Name,
		module.Position,
		module.IsPusblished,
		module.UnlockDate,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("moduleRepository.Create: %w", err)
	}

	return id, nil
}

func (r *moduleRepository) Delete(ctx context.Context, moduleID string) error {
	query := `DELETE FROM modules WHERE id = $1`

	result, err := r.db.Exec(ctx, query, moduleID)
	if err != nil {
		return fmt.Errorf("moduleRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("moduleRepository.Delete: module not found with ID: %s", moduleID)
	}
	return nil
}

func (r *moduleRepository) GetCourseModules(ctx context.Context, courseID string) ([]*domain.Module, error) {
	query := `
	SELECT id, course_id, name, position, is_published, unlock_date, created_at,updated_at 
	FROM modules WHERE course_id = $1`

	rows, err := r.db.Query(ctx, query, courseID)
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.GetCourseModules: %w", err)
	}

	modules, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Module])
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.GetCourseModules: %w", err)
	}

	return modules, nil
}
