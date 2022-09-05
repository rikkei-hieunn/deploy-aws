package repository

import (
	"chikuseki-check/infrastructure"
	"chikuseki-check/infrastructure/db/rdb"
	"chikuseki-check/model"
	"chikuseki-check/pkg/utils"
	"context"
	"database/sql"
	"fmt"
)

// tickDBRepository provides functions to import data into sorting database
type tickDBRepository struct {
	db rdb.ITickDBHandler
}

// NewTickDBRepository constructor
func NewTickDBRepository(infra *infrastructure.Infra) ITickDBRepository {
	return &tickDBRepository{
		infra.TickDB,
	}
}

// InitConnection init database connection
func (r tickDBRepository) InitConnection(ctx context.Context, endpoint, dbName, kei, dataType string) error {
	return r.db.InitConnection(endpoint, dbName, kei, dataType)
}

// CheckTableExists check table exists in database
func (r tickDBRepository) CheckTableExists(ctx context.Context, prefix, zxd, kubun, hassin, dbName, kei, dataType string) bool {
	// get table name from prefix, kubun, hassin, zxd
	var tableName string
	kubunInsteadOf, isExists := model.KubunsInsteadOf[kubun]
	if !isExists {
		tableName = fmt.Sprintf(model.TableNameFormat, prefix, kubun, hassin, zxd)
	} else {
		tableName = fmt.Sprintf(model.TableNameFormat, prefix, kubunInsteadOf, hassin, zxd)
	}

	// query check table exists
	queryString := fmt.Sprintf(model.QueryCheckTableExists, dbName, tableName)
	connectionKey := kei + model.StrokeCharacter + dataType
	result, err := r.db.Query(ctx, queryString, connectionKey, parseTableName)
	if err != nil {
		return false
	}

	// if table name result does not equal table name input then return false
	resultString, ok := result.(string)
	if !ok || resultString != tableName {
		return false
	}

	return true
}

// CountNumberRecords count total number records
func (r tickDBRepository) CountNumberRecords(ctx context.Context, prefix, zxd, kubun, hassin, dbName, kei, dataType string) (*int, error) {
	// get table name from prefix, kubun, hassin, zxd
	var tableName string
	kubunInsteadOf, isExists := model.KubunsInsteadOf[kubun]
	if !isExists {
		tableName = fmt.Sprintf(model.TableNameFormat, prefix, kubun, hassin, zxd)
	} else {
		tableName = fmt.Sprintf(model.TableNameFormat, prefix, kubunInsteadOf, hassin, zxd)
	}

	// query count number records
	queryString := fmt.Sprintf(model.QueryCountAllTable, tableName)
	connectionKey := kei + model.StrokeCharacter + dataType
	result, err := r.db.Query(ctx, queryString, connectionKey, parseNumberRecords)
	if err != nil {
		return nil, err
	}

	// if table name result does not equal table name input then return false
	resultString, ok := result.(int)
	if !ok {
		return nil, fmt.Errorf("get total number records fail")
	}

	return &resultString, nil
}

// parseTableName parse from row to table name string
func parseTableName(rows *sql.Rows) (interface{}, error) {
	for rows.Next() {
		var tableName interface{}
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		if tableName != nil {
			return utils.ToString(tableName), nil
		}
	}

	return nil, fmt.Errorf("empty data")
}

// parseNumberRecords parse from row to total records
func parseNumberRecords(rows *sql.Rows) (interface{}, error) {
	for rows.Next() {
		var total int
		err := rows.Scan(&total)
		if err != nil {
			return nil, err
		}

		return total, nil
	}

	return nil, fmt.Errorf("empty data")
}
