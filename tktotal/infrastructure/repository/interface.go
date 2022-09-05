package repository

import (
	"bytes"
	"context"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(ctx context.Context, path string) (*bytes.Buffer, error)
	Upload(ctx context.Context, path string, body bytes.Buffer) error
	GetObjectKeys(ctx context.Context, path string) ([]string, error)
}