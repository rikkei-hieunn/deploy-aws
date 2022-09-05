/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"recreate-one-minute/configs"
	"recreate-one-minute/model"
	"strings"

	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	clients map[string]*sql.DB
}

// NewTickDBHandler constructor
func NewTickDBHandler(config *configs.TickDB) (ITickDBHandler, error) {
	mysqlClient := &mysql{}
	clients := make(map[string]*sql.DB)
	for endpoint, kubunHasshinPairs := range config.Endpoints {
		temp := strings.Split(endpoint, model.StrokeCharacter)
		dbEndpoint := temp[0]
		dbName := temp[1]
		client, err := mysqlClient.init(config, dbEndpoint, dbName)
		if err != nil {
			return nil, fmt.Errorf("enpoint :%s,db name :%s :  init db fail : %w", dbEndpoint, dbName, err)
		}
		for _, kubunHasshin := range kubunHasshinPairs {
			clients[kubunHasshin] = client
		}
	}
	mysqlClient.clients = clients

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

	if pingErr := client.Ping(); pingErr != nil {
		return nil, pingErr
	}

	client.SetMaxOpenConns(config.MaxOpenConnection)
	client.SetMaxIdleConns(config.MaxIdleConnection)

	return client, nil
}

//Close DB
func (c *mysql) Close() error {
	for _, client := range c.clients {
		if err := client.Close(); err != nil {
			return err
		}
	}

	return nil
}

//ExecuteQuery execute query
func (c *mysql) ExecuteQuery(ctx context.Context, query string, key string, parsingObject ParseSQLObject) ([]interface{}, error) {
	client, isExisted := c.clients[key]
	if !isExisted {
		return nil, fmt.Errorf("client not found with key : %s", key)
	}
	rows, err := client.Query(query)
	if err != nil {
		return nil, err
	}

	return parsingObject(rows), nil
}
