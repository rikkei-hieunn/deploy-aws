/*
Package repository implements logics repository.
*/
package repository

import (
	"bytes"
	"context"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}

// ITickDBRepository provides sorting db interfaces
type ITickDBRepository interface {
	InitConnection(ctx context.Context, endpoint, dbName, kei, dataType string) error
	CheckTableExists(ctx context.Context, prefix, zxd, kubun, hassin, dbName, kei, dataType string) bool
	CountNumberRecords(ctx context.Context, prefix, zxd, kubun, hassin, dbName, kei, dataType string) (*int, error)
}
