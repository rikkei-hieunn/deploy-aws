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
	ExecuteQuery(ctx context.Context, query, key string, parsingObject ParseSQLObject) ([]interface{}, error)
	Close() error
}

//ParseSQLObject parse sql rows to data
type ParseSQLObject func(rows *sql.Rows) []interface{}
