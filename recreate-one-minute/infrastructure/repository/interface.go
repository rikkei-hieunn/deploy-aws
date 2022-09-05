/*
Package repository provides logics implementation
*/
package repository

import (
	"bytes"
	"context"
	"recreate-one-minute/model"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}

// ITickDBRepository Structure of interface DB
type ITickDBRepository interface {
	GetDataFromCandleManagement(ctx context.Context, key string, args []interface{}, validRecord *[]model.Record) error
	Close() error
}
