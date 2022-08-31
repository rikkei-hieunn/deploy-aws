package repository

import (
	"context"
	"create-table/configs"
	"create-table/infrastructure"
	"create-table/infrastructure/db/rdb"
	"create-table/model"
	"create-table/pkg/utils"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
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
func (r tickDBRepository) InitConnection(ctx context.Context, endpoint, dbName string) error {
	return r.db.InitConnection(endpoint, dbName)
}

// CheckTableExists check table exists in database
func (r tickDBRepository) CheckTableExists(ctx context.Context, targetCreateTable configs.TargetCreateTable, zxd string) bool {
	// get table name from prefix, kubun, hassin, zxd
	var tableName string
	kubunInsteadOf, isExists := model.KubunsInsteadOf[targetCreateTable.QKbn]
	if !isExists {
		tableName = fmt.Sprintf(model.TableNameFormat, targetCreateTable.TablePrefix, targetCreateTable.QKbn, targetCreateTable.Sndc, zxd)
	} else {
		tableName = fmt.Sprintf(model.TableNameFormat, targetCreateTable.TablePrefix, kubunInsteadOf, targetCreateTable.Sndc, zxd)
	}

	// query check table exists
	queryString := fmt.Sprintf(model.QueryCheckTableExists, targetCreateTable.DBName, tableName)
	connectionKey := targetCreateTable.Endpoint + model.StrokeCharacter + targetCreateTable.DBName
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

// CreateTable generate query and execute create table
func (r tickDBRepository) CreateTable(ctx context.Context, targetCreateTable configs.TargetCreateTable, zxd string) (string, error) {
	// get table name from prefix, kubun, hassin, zxd
	var tableName string
	kubunInsteadOf, isExists := model.KubunsInsteadOf[targetCreateTable.QKbn]
	if !isExists {
		tableName = fmt.Sprintf(model.TableNameFormat, targetCreateTable.TablePrefix, targetCreateTable.QKbn, targetCreateTable.Sndc, zxd)
	} else {
		tableName = fmt.Sprintf(model.TableNameFormat, targetCreateTable.TablePrefix, kubunInsteadOf, targetCreateTable.Sndc, zxd)
	}

	var queryString string
	if targetCreateTable.DataType == model.OneMinuteData {
		queryString = genCreateTableOneMinuteData(targetCreateTable, tableName)
	} else if targetCreateTable.DataType == model.TickDataString ||
		targetCreateTable.DataType == model.KehaiDataString || targetCreateTable.DataType == model.JishouDataString {
		queryString = genCreateTableMainData(targetCreateTable, tableName)
	} else {
		queryString = genCreateTableExtendData(targetCreateTable, tableName)
	}

	if queryString == model.EmptyString {
		return tableName, nil
	}
	connectionKey := targetCreateTable.Endpoint + model.StrokeCharacter + targetCreateTable.DBName

	return tableName, r.db.Execute(ctx, queryString, connectionKey)
}

// genCreateTableMainData generate query string create table for main data
func genCreateTableMainData(targetCreateTable configs.TargetCreateTable, tableName string) string {
	table, isExists := model.Tables[targetCreateTable.QKbn][targetCreateTable.DataType]
	if !isExists {
		return model.EmptyString
	}

	sbValues := strings.Builder{}
	for index := range table.Elements {
		if table.Elements[index].Column == "QCD" || table.Elements[index].Column == "TIME" ||
			table.Elements[index].Column == "TKZXD" || table.Elements[index].Column == "TKTIM" ||
			table.Elements[index].Column == "TKSERIALNUMBER" || table.Elements[index].Column == "TKQKBN" ||
			table.Elements[index].Column == "SNDC" {
			continue
		}

		columnName := table.Elements[index].Column
		if columnName == "NOT" {
			columnName = "`" + columnName + "`"
		}

		sbValues.WriteString(columnName)
		sbValues.WriteString(" varchar(")
		sbValues.WriteString(strconv.Itoa(table.Elements[index].Length))
		sbValues.WriteString("),")
	}

	return fmt.Sprintf(model.QueryCreateTableForMainData, tableName, sbValues.String())
}

// genCreateTableExtendData generate query string create table for extend data
func genCreateTableExtendData(targetCreateTable configs.TargetCreateTable, tableName string) string {
	table, isExists := model.Tables[targetCreateTable.QKbn][targetCreateTable.DataType]
	if !isExists {
		return model.EmptyString
	}

	sbValues := strings.Builder{}
	for index := range table.Elements {
		if table.Elements[index].Column == "QCD" || table.Elements[index].Column == "TIME" ||
			table.Elements[index].Column == "HTKZXD" || table.Elements[index].Column == "HTKTIM" ||
			table.Elements[index].Column == "TKSERIALNUMBER" || table.Elements[index].Column == "TKQKBN" ||
			table.Elements[index].Column == "SNDC" {
			continue
		}

		columnName := table.Elements[index].Column
		if columnName == "NOT" {
			columnName = "`" + columnName + "`"
		}

		sbValues.WriteString(columnName)
		sbValues.WriteString(" varchar(")
		sbValues.WriteString(strconv.Itoa(table.Elements[index].Length))
		sbValues.WriteString("),")
	}

	return fmt.Sprintf(model.QueryCreateTableForExtendData, tableName, sbValues.String())
}

// genCreateTableOneMinuteData generate query string create table for one minute data
func genCreateTableOneMinuteData(targetCreateTable configs.TargetCreateTable, tableName string) string {
	table, isExists := model.Tables[targetCreateTable.QKbn][targetCreateTable.DataType]
	if !isExists {
		return model.EmptyString
	}

	sbValues := strings.Builder{}
	for index := range table.Elements {
		if table.Elements[index].Column == "QCD" || table.Elements[index].Column == "TIME" ||
			table.Elements[index].Column == "TKQKBN" || table.Elements[index].Column == "SNDC" {
			continue
		}

		columnName := table.Elements[index].Column
		if columnName == "NOT" {
			columnName = "`" + columnName + "`"
		}

		sbValues.WriteString(columnName)
		sbValues.WriteString(" varchar(")
		sbValues.WriteString(strconv.Itoa(table.Elements[index].Length))
		sbValues.WriteString("),")
	}

	return fmt.Sprintf(model.QueryCreateTableForOneMinuteData, tableName, sbValues.String())
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
