/*
Package rdb implements logics about database.
*/
package rdb

import (
	"context"
	"data-del/configs"
	"data-del/model"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"

	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

//mysql define sql clients object  for both tick and kehai
type mysql struct {
	clients map[string]map[string]*sql.DB
}

// NewTickDBHandler constructor
func NewTickDBHandler(config *configs.TickDB) (ITickDBHandler, error) {
	mysqlClient := &mysql{}
	mapClients := make(map[string]map[string]*sql.DB)
	for dbType, mapEndpoints := range config.Endpoints {
		clients := make(map[string]*sql.DB)
		for endpoint, kubunHasshinPairs := range mapEndpoints {
			temp := strings.Split(endpoint, model.StrokeCharacter)
			dbEndpoint := temp[0]
			dbName := temp[1]
			client, err := mysqlClient.init(config, dbEndpoint, dbName)
			if err != nil {
				return nil, fmt.Errorf("enpoint :%s,db name :%s :  init db fail : %w", dbEndpoint, dbName, err)
			}
			for i := range kubunHasshinPairs {
				clients[kubunHasshinPairs[i]] = client
			}
		}
		mapClients[dbType] = clients
	}
	mysqlClient.clients = mapClients

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
		dbName)
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

//Close disconnection all database for tick and kehai
func (c *mysql) Close() error {
	for _, client := range c.clients {
		for _, db := range client {
			if err := db.Close(); err != nil {
				return err
			}
		}
	}

	return nil
}

//Exec execute sql not return data
func (c *mysql) Exec(ctx context.Context, sql, dbType string) error {
	client, isExisted := c.clients[dbType]
	if !isExisted {
		return fmt.Errorf("clients is not found with data-type %s ", dbType)
	}
	for _, db := range client {
		_, err := db.ExecContext(ctx, sql)
		if err != nil {
			return fmt.Errorf("exec query fail, sql: %s ", sql)
		}
	}

	return nil
}

//Query execute sql and return data
func (c *mysql) Query(ctx context.Context, sql, dbType string, parse ParseObject) (interface{}, error) {
	var result []string
	temp := make(map[string]interface{})
	client, isExisted := c.clients[dbType]
	if !isExisted {
		return nil, fmt.Errorf("clients is not found with data-type %s ", dbType)
	}
	for _, db := range client {
		rows, err := db.Query(sql)
		if err != nil {
			return nil, err
		}
		tables := parse(rows)
		for _, table := range tables {
			//TODO try to not casting
			tableString, ok := table.(string)
			if !ok {
				log.Error().Msg("casting error")
			}
			temp[tableString] = nil
		}
	}
	for tableName := range temp {
		result = append(result, tableName)
	}

	return result, nil
}
