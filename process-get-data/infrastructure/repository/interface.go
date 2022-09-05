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
	InitConnection(ctx context.Context, host, dbName, kubun, hassin string) error
	InsertData(ctx context.Context, sql, kubun, hassin string, args []interface{}) error
}
