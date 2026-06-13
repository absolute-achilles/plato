package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ModuleContentRepository interface {
	Get(ctx context.Context, moduleContentID string) (*domain.ModuleContent, error)
	Create(ctx context.Context, moduleContent *domain.ModuleContent) (id string, err error)
	Delete(ctx context.Context, moduleContentID string) error

	GetModuleContents(ctx context.Context, moduleID string) ([]*domain.ModuleContent, error)
}

type moduleContentRepository struct {
	db *pgxpool.Pool
}

func NewModuleContentRepository(db *pgxpool.Pool) ModuleContentRepository {
	return &moduleContentRepository{db: db}
}

func (r *moduleContentRepository) Get(ctx context.Context, moduleContentID string) (*domain.ModuleContent, error) {
	query := `
	SELECT id, module_id, title, type, body_content, position, is_published,created_at, updated_at 
	FROM module_contents WHERE id = $1`

	rows, err := r.db.Query(ctx, query, moduleContentID)
	if err != nil {
		return nil, fmt.Errorf("moduleContentRepository.Get: %w", err)
	}
	module, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.ModuleContent])
	if err != nil {
		return nil, fmt.Errorf("moduleContentRepository.Get: %w", err)
	}

	return module, nil
}

func (r *moduleContentRepository) Create(ctx context.Context, moduleContent *domain.ModuleContent) (id string, err error) {
	if moduleContent == nil {
		return "", fmt.Errorf("moduleContentRepository.Create: Empty Module Request")
	}

	if moduleContent.ModuleID == "" {
		return "", fmt.Errorf("moduleContentRepository.Create: Empty Course ID")
	}

	insertQuery := `
		INSERT INTO module_contents (module_id, title, type, body_content, position, is_published)
		VALUES ($1, $2, $3, $4, $5 , $6)
		RETURNING id
	`

	if err := r.db.QueryRow(
		ctx,
		insertQuery,
		moduleContent.ModuleID,
		moduleContent.Title,
		moduleContent.ModuleContentType,
		moduleContent.BodyContent,
		moduleContent.Position,
		moduleContent.IsPublished,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("moduleRepository.Create: %w", err)
	}

	return id, nil
}

func (r *moduleContentRepository) Delete(ctx context.Context, moduleContentID string) error {
	query := `DELETE FROM module_contents WHERE id = $1`

	result, err := r.db.Exec(ctx, query, moduleContentID)
	if err != nil {
		return fmt.Errorf("moduleContentRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("moduleContentRepository.Delete: module content not found with ID: %s", moduleContentID)
	}
	return nil
}

func (r *moduleContentRepository) GetModuleContents(ctx context.Context, moduleID string) ([]*domain.ModuleContent, error) {
	query := `
	SELECT id, module_id, title, type, body_content, position, is_published,created_at, updated_at 
	FROM module_contents WHERE module_id = $1`

	rows, err := r.db.Query(ctx, query, moduleID)
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.GetCourseModules: %w", err)
	}

	modules, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.ModuleContent])
	if err != nil {
		return nil, fmt.Errorf("moduleRepository.GetCourseModules: %w", err)
	}

	return modules, nil
}
