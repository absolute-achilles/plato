package storage

import (
	"context"
	"io"
)

type Storage interface {
	// Upload streams a file to the bucket.
	// Returns the generated object key/filename and any error.
	// TODO: use chan instead to stream the bytes that have been uploaded to have progression in UX?
	Upload(ctx context.Context, filename string, content io.Reader, size int64, contentType string) (string, error)

	// Delete removes a file from the bucket using its object key.
	Delete(ctx context.Context, filename string) error

	// GetUrl generates a public or pre-signed URL for the given file.
	GetUrl(ctx context.Context, filename string) (string, error)
}
