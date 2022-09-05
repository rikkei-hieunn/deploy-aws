/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
	"database/sql"
)

// ITickDBHandler sorting database interfaces
type ITickDBHandler interface {
	Query(ctx context.Context, sql, dbType string, parse ParseObject) (interface{}, error)
	Exec(ctx context.Context, sql, dbType string) error
	Close() error
}

//ParseObject parse SQL object
type ParseObject func(rows *sql.Rows) []interface{}
