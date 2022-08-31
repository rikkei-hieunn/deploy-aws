/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
)

// ITickDBHandler sorting database interfaces
type ITickDBHandler interface {
	InitTx(ctx context.Context, kei string) error
	ExecWithTx(ctx context.Context, sql string, kei string, args []interface{}) error
	RollbackTx(kei string) error
	CommitTx(kei string) error
	Close() error
}
