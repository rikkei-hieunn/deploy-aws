package repository

import (
	"context"
	"data-del/infrastructure"
	"data-del/infrastructure/db/rdb"
	"data-del/model"
	"database/sql"
	"fmt"
)

const (
	baseSelectTableSQL         = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME LIKE '%s_%s_%s_%s';"
	baseSelectTableSQLOnlyDate = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME LIKE '%s_%s';"
	baseSQLDropTable           = "DROP TABLE IF EXISTS %s ;"
)

// tickDBRepository Structure of repository DB
type tickDBRepository struct {
	db rdb.ITickDBHandler
}

// NewTickDBRepository Initialize a Repository DB
func NewTickDBRepository(infra *infrastructure.Infra) ITickDBRepository {
	return &tickDBRepository{
		infra.TickDB,
	}
}

//GetListAvailableTable get a list unique table
func (t *tickDBRepository) GetListAvailableTable(ctx context.Context, kubun, hasshin, date, dbType string) ([]string, error) {
	_sql := fmt.Sprintf(baseSelectTableSQL, model.PercentSign, kubun, hasshin, date)
	result, err := t.db.Query(ctx, _sql, dbType, t.parseTable)
	if err != nil {
		return nil, err
	}

	tables, ok := result.([]string)
	if !ok {
		return nil, fmt.Errorf("cast table fail")
	}

	return tables, nil
}

//GetListAvailableTableWithSameDate select all table same date
func (t *tickDBRepository) GetListAvailableTableWithSameDate(ctx context.Context, date, dbType string) ([]string, error) {
	_sql := fmt.Sprintf(baseSelectTableSQLOnlyDate, model.PercentSign, date)
	result, err := t.db.Query(ctx, _sql, dbType, t.parseTable)
	if err != nil {
		return nil, err
	}
	tables, ok := result.([]string)
	if !ok {
		return nil, fmt.Errorf("cast table fail")
	}

	return tables, nil
}

//DropTable drop table
func (t *tickDBRepository) DropTable(ctx context.Context, tableName, dataType string) error {
	_sql := fmt.Sprintf(baseSQLDropTable, tableName)

	return t.db.Exec(ctx, _sql, dataType)
}

//Close disconnect db
func (t *tickDBRepository) Close() error {
	return t.db.Close()
}

func (t *tickDBRepository) parseTable(row *sql.Rows) []interface{} {
	var results []interface{}
	var tableName string
	for row.Next() {
		err := row.Scan(&tableName)
		if err != nil {
			return nil
		}
		results = append(results, tableName)
	}

	return results
}
