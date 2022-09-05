package repository

import (
	"bytes"
	"insert-calendar-infos/infrastructure"
	"insert-calendar-infos/infrastructure/s3"
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

// Download download file
func (s *s3Repository) Download(path string) ([]byte, error) {
	return s.s3Client.GetObject(path)
}

// Upload upload file
func (s *s3Repository) Upload(prefix string, data bytes.Buffer) error {
	return s.s3Client.PutObject(prefix, data)
}