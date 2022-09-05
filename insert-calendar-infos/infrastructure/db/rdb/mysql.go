/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/model"
	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	clients map[string]*sql.DB
	txs     map[string]*sql.Tx
}

// NewTickDBHandler constructor init connection database
func NewTickDBHandler(config *configs.TickDB) (ITickDBHandler, error) {
	mysqlClient := &mysql{
		clients: make(map[string]*sql.DB),
		txs:     make(map[string]*sql.Tx),
	}
	connectionFirstKei, err := mysqlClient.init(config, config.CalendarEndpoint.DBKei1Endpoint, config.CalendarEndpoint.DBKei1Name)
	if err != nil {
		return nil, err
	}

	connectionSecondKei, err := mysqlClient.init(config, config.CalendarEndpoint.DBKei2Endpoint, config.CalendarEndpoint.DBKei2Name)
	if err != nil {
		return nil, err
	}
	mysqlClient.clients[model.TheFirstKei] = connectionFirstKei
	mysqlClient.clients[model.TheSecondKei] = connectionSecondKei

	return mysqlClient, nil
}

// init initialize tickDB
func (c *mysql) init(config *configs.TickDB, host, dbName string) (*sql.DB, error) {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		host,
		config.Port,
		dbName,
	)

	client, err := sql.Open(config.DriverName, connectInfo)
	if err != nil {
		return nil, err
	}

	pingErr := client.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	client.SetMaxOpenConns(config.MaxOpenConnection)
	client.SetMaxIdleConns(config.MaxIdleConnection)

	return client, nil
}

// InitTx init transaction
func (c *mysql) InitTx(ctx context.Context, kei string) error {
	currentClient, isExists := c.clients[kei]
	if !isExists {
		return fmt.Errorf("database connection not found")
	}

	tx, err := currentClient.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	c.txs[kei] = tx

	return nil
}

// ExecWithTx execute query in transaction
func (c *mysql) ExecWithTx(ctx context.Context, sql string, kei string, args []interface{}) error {
	transaction, isExists := c.txs[kei]
	if !isExists {
		return fmt.Errorf("transaction not found")
	}
	_, err := transaction.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// RollbackTx rollback transaction
func (c *mysql) RollbackTx(kei string) error {
	transaction, isExists := c.txs[kei]
	if !isExists {
		return fmt.Errorf("transaction not found")
	}
	err := transaction.Rollback()
	if err != nil {
		return err
	}

	return nil
}

// CommitTx commit transaction
func (c *mysql) CommitTx(kei string) error {
	transaction, isExists := c.txs[kei]
	if !isExists {
		return fmt.Errorf("transaction not found")
	}
	err := transaction.Commit()
	if err != nil {
		return err
	}

	return err
}

// Close closed connection
func (c *mysql) Close() error {
	for _, cli := range c.clients {
		if err := cli.Close(); err != nil {
			return err
		}
	}

	return nil
}
