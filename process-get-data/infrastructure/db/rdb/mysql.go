package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"process-get-data/configs"
	"process-get-data/model"

	// register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	config  *configs.TickDB
	clients map[string]*sql.DB
}

// NewTickDBHandler constructor
func NewTickDBHandler(cfg *configs.TickDB) (ITickDBHandler, error) {
	return &mysql{
		config:  cfg,
		clients: make(map[string]*sql.DB),
	}, nil
}

// InitConnection init db connection
func (c *mysql) InitConnection(host, dbName, kubun, hassin string) error {
	connectionKey := kubun + model.StrokeCharacter + hassin
	_, isExists := c.clients[connectionKey]
	if isExists {
		return nil
	}
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.config.User,
		c.config.Password,
		host,
		c.config.Port,
		dbName,
	)

	client, err := sql.Open(c.config.DriverName, connectInfo)
	if err != nil {
		return err
	}

	if pingErr := client.Ping(); pingErr != nil {
		return pingErr
	}

	client.SetMaxOpenConns(c.config.MaxOpenConnection)
	client.SetMaxIdleConns(c.config.MaxIdleConnection)
	c.clients[connectionKey] = client

	return nil
}

// Execute execute sql insert many
func (c *mysql) Execute(ctx context.Context, sql string, key string, args []interface{}) error {
	currentClient, isExists := c.clients[key]
	if !isExists {
		return fmt.Errorf("database connection not found")
	}
	_, err := currentClient.ExecContext(ctx, sql, args...)

	return err
}

// Query query data from database
func (c *mysql) Query(ctx context.Context, sql, key string, rowsHandler ParseStructure) (interface{}, error) {
	client, isExisted := c.clients[key]
	if !isExisted {
		return nil, fmt.Errorf("database connection not found")
	}
	rows, err := client.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	result, err := rowsHandler(rows)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	return result, nil
}

// Close remove client connection
func (c *mysql) Close() error {
	for _, client := range c.clients {
		if err := client.Close(); err != nil {
			return err
		}
	}

	return nil
}
