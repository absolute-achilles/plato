package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContentAttachmentRepository interface {
	Get(ctx context.Context, id string) (*domain.ContentAttachment, error)
	Create(ctx context.Context, contentAttachment *domain.ContentAttachment) (id string, err error)
	Delete(ctx context.Context, id string) error

	GetModuleContentAttachments(ctx context.Context, moduleContentID string) ([]*domain.ContentAttachment, error)
}

type contentAttachmentRepository struct {
	db *pgxpool.Pool
}

func NewContentAttachmentRepository(db *pgxpool.Pool) ContentAttachmentRepository {
	return &contentAttachmentRepository{db: db}
}

func (r *contentAttachmentRepository) Get(ctx context.Context, id string) (*domain.ContentAttachment, error) {
	query := `
	SELECT id, module_content_id, name, url, size_bytes, type, created_at 
	FROM content_attachments WHERE id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("contentAttachmentRepository.Get: %w", err)
	}
	module, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.ContentAttachment])
	if err != nil {
		return nil, fmt.Errorf("contentAttachmentRepository.Get: %w", err)
	}

	return module, nil
}

func (r *contentAttachmentRepository) Create(ctx context.Context, contentAttachment *domain.ContentAttachment) (id string, err error) {
	if contentAttachment == nil {
		return "", fmt.Errorf("contentAttachmentRepository.Create: Empty Content Request")
	}

	if contentAttachment.ModuleContentID == "" {
		return "", fmt.Errorf("contentAttachmentRepository.Create: Empty Module Content ID")
	}

	insertQuery := `
		INSERT INTO content_attachments (module_content_id, name, url, size_bytes, type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	if err := r.db.QueryRow(
		ctx,
		insertQuery,
		contentAttachment.ModuleContentID,
		contentAttachment.Name,
		contentAttachment.URL,
		contentAttachment.SizeBytes,
		contentAttachment.FileType,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("contentAttachmentRepository.Create: %w", err)
	}

	return id, nil
}

func (r *contentAttachmentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM content_attachments WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("contentAttachmentRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("contentAttachmentRepository.Delete: content attachment not found with ID: %s", id)
	}
	return nil
}

func (r *contentAttachmentRepository) GetModuleContentAttachments(ctx context.Context, moduleContentID string) ([]*domain.ContentAttachment, error) {
	query := `
    SELECT id, module_content_id, name, url, size_bytes, type, created_at 
    FROM content_attachments WHERE module_content_id = $1`

	rows, err := r.db.Query(ctx, query, moduleContentID)
	if err != nil {
		return nil, fmt.Errorf("contentAttachmentRepository.GetModuleContentAttachments: %w", err)
	}

	contentAttachments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.ContentAttachment])
	if err != nil {
		return nil, fmt.Errorf("contentAttachmentRepository.GetModuleContentAttachments: %w", err)
	}

	return contentAttachments, nil
}
