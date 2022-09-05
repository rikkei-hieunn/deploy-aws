// Package repository for s3
package repository

import (
	"bytes"
	"context"
	"tktotal/infrastructure"
	"tktotal/infrastructure/storage/s3"
)

type s3Repository struct {
	s3Client s3.IS3Handler
}

// NewStorageRepository storage repository constructor
func NewStorageRepository(infra *infrastructure.Infra) IS3Repository {
	return &s3Repository{
		s3Client: infra.S3Handler,
	}
}

// Download func download file
func (s *s3Repository) Download(ctx context.Context, path string) (*bytes.Buffer, error) {
	return s.s3Client.GetObject(ctx, path)
}

func (s *s3Repository) Upload(ctx context.Context, path string, body bytes.Buffer) error {
	return s.s3Client.PutObject(ctx, path, body)
}

func (s *s3Repository) GetObjectKeys(ctx context.Context, path string) ([]string, error) {
	return s.s3Client.GetObjects(ctx, path)
}
