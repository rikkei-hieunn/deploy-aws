/*
Package repository implements logics repository.
*/
package repository

import (
	"bytes"
	"context"
	"create-table/configs"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}

// ITickDBRepository provides sorting db interfaces
type ITickDBRepository interface {
	InitConnection(ctx context.Context, endpoint, dbName string) error
	CheckTableExists(ctx context.Context, targetCreateTable configs.TargetCreateTable, zxd string) bool
	CreateTable(ctx context.Context, targetCreateTable configs.TargetCreateTable, zxd string) (string, error)
}
