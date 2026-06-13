package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/absolute-achilles/plato/internal/storage"
)

type ContentAttachmentService interface {
	UploadAttachment(ctx context.Context, req *domain.ContentAttachment, file io.Reader) (string, error)
	DeleteAttachment(ctx context.Context, id string) error
}

type contentAttachmentService struct {
	repo    repository.ContentAttachmentRepository
	storage storage.Storage
}

func NewContentAttachmentService(repo repository.ContentAttachmentRepository, storage storage.Storage) ContentAttachmentService {
	return &contentAttachmentService{
		repo:    repo,
		storage: storage,
	}
}

func (s *contentAttachmentService) UploadAttachment(ctx context.Context, attachment *domain.ContentAttachment, file io.Reader) (string, error) {
	fileURL, err := s.storage.Upload(ctx, attachment.Name, file, int64(attachment.SizeBytes), attachment.FileType)
	if err != nil {
		return "", fmt.Errorf("failed to upload to storage: %w", err)
	}

	attachment.URL = fileURL

	id, err := s.repo.Create(ctx, attachment)
	if err != nil {
		// If DB fails, attempt to delete the orphaned file from storage
		_ = s.storage.Delete(ctx, attachment.Name)
		return "", fmt.Errorf("failed to save attachment record to database: %w", err)
	}

	return id, nil
}

func (s *contentAttachmentService) DeleteAttachment(ctx context.Context, id string) error {
	attachment, err := s.repo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("attachment not found: %w", err)
	}

	// If this fails, the file remains in storage (safe).
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete from database: %w", err)
	}

	// If this fails, we have an orphaned file in the bucket, but the app state is consistent.
	if err := s.storage.Delete(ctx, attachment.Name); err != nil {
		slog.Warn("failed to delete file from storage", "Err", err)
	}

	return nil
}
