package storage

import (
	"context"
	"io"

	"github.com/absolute-achilles/plato/internal/domain"
)

type Storage interface {
	Upload(ctx context.Context, filename string, content io.Reader, contentType domain.ContentType) (string, error)
	Delete(ctx context.Context, filename string)
	GetUrl(filename string) string
}
