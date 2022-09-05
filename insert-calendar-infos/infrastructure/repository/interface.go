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

// ITickDBRepository Structure of interface DB
type ITickDBRepository interface {
	CommitTx(kei string)
	RollbackTx(kei string)
	InitTx(ctx context.Context, kei string) error
	ExecWithTx(ctx context.Context, sql, kei string, args []interface{}) error
	Close() error
}

// IFilebusRepository Structure of interface filebus repository
type IFilebusRepository interface {
	DownloadFile(path, file string) ([]byte, error)
}
