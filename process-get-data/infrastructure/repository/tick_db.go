package repository

import (
	"context"
	"process-get-data/infrastructure"
	"process-get-data/infrastructure/db/rdb"
	"process-get-data/model"
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
func (r tickDBRepository) InitConnection(ctx context.Context, host, dbName, kubun, hassin string) error {
	return r.db.InitConnection(host, dbName, kubun, hassin)
}

// InsertData insert data
func (r tickDBRepository) InsertData(ctx context.Context, sql, kubun, hassin string, args []interface{}) error {
	connectionKey := kubun + model.StrokeCharacter + hassin

	return r.db.Execute(ctx, sql, connectionKey, args)
}
