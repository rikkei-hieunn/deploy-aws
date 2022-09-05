/*
Package createtable implements logics create table.
*/
package createtable

import (
	"context"
	"create-table/configs"
)

// ITableCreator define method create table
type ITableCreator interface {
	CreateTables(ctx context.Context, quoteCodes []configs.TargetCreateTable) []error
}
