/*
Package countdata implements logics count data from database.
*/
package countdata

import (
	"context"
)

// ITableCreator define method create table
type ITableCreator interface {
	CountData(ctx context.Context) error
}
