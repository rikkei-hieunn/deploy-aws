/*
Package repository define all repository services
*/
package repository

import (
	"bytes"
	"context"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	GetObject(ctx context.Context, path string) (*bytes.Buffer, error)
	CheckObjectExists(ctx context.Context, Prefix string) (bool, error)
}

// ITickDBRepository Structure of interface DB
type ITickDBRepository interface {
	DropTable(ctx context.Context, tableName, dbType string) error
	GetListAvailableTable(ctx context.Context, kubun, hasshin, date, dbType string) ([]string, error)
	GetListAvailableTableWithSameDate(ctx context.Context,date, dbType string) ([]string, error)
	Close() error
}
