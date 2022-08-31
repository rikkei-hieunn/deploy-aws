/*
Package insertdata implements logics insert data and create cron tab.
*/
package insertdata

import (
	"context"
	"process-get-data/configs"
)

// ITableCreator define method create table
type ITableCreator interface {
	InsertData(ctx context.Context, quoteCodes map[string]configs.QuoteCodes, kei string) []error
}
