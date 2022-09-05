package repository

import (
	"bytes"
	"context"
	"data-del/infrastructure"
	"data-del/infrastructure/storage/s3"
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

// GetObject get file in s3
func (s *s3Repository) GetObject(ctx context.Context, path string) (*bytes.Buffer, error) {
	return s.s3Client.GetObject(ctx, path)
}

// CheckObjectExists check file exists in folder
func (s *s3Repository) CheckObjectExists(ctx context.Context, prefix string) (bool, error) {
	return s.s3Client.CheckObjectExists(ctx, prefix)
}
