/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
	"database/sql"
)

//ParseStructure parse record from db to struct
type ParseStructure func(rows *sql.Rows) (interface{}, error)

// ITickDBHandler sorting database interfaces
type ITickDBHandler interface {
	Query(ctx context.Context, sql, key string, rowsHandler ParseStructure) (interface{}, error)
	Execute(ctx context.Context, sql string, key string, args []interface{}) error
	InitConnection(host, dbName, kubun, hassin string) error
	Close() error
}
