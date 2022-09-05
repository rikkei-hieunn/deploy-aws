package s3

import (
	"bytes"
	"context"
)

// IS3Handler interface S3 handler
type IS3Handler interface {
	GetObject(ctx context.Context, path string) (*bytes.Buffer, error)
	PutObject(ctx context.Context, path string, body bytes.Buffer) error
	GetObjects(ctx context.Context, path string) ([]string, error)
}
