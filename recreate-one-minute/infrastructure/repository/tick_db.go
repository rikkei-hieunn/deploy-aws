package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"recreate-one-minute/infrastructure"
	"recreate-one-minute/infrastructure/db/rdb"
	"recreate-one-minute/model"
	"recreate-one-minute/pkg"
)

const baseSelectSQL = "SELECT OPERATION_TYPE,TKQKBN,SNDC,CREATE_DAY,CREATE_TIME,START_INDEX,FOLDER_PATH FROM %s WHERE CREATE_STATUS ='%s'"

// TickDBRepository Structure of repository DB
type tickDBRepository struct {
	db rdb.ITickDBHandler
}

// NewTickDBRepository Initialize a Repository DB
func NewTickDBRepository(infra *infrastructure.Infra) ITickDBRepository {
	return &tickDBRepository{
		infra.TickDB,
	}
}

//GetDataFromCandleManagement execute query
func (s *tickDBRepository) GetDataFromCandleManagement(ctx context.Context, key string, args []interface{}, validRecords *[]model.Record) error {
	query := fmt.Sprintf(baseSelectSQL, args...)
	data, err := s.db.ExecuteQuery(ctx, query, key, parseCandleRecords)
	if err != nil {
		return fmt.Errorf("execute sql : %s error %w ", query, err)
	}
	for i := range data {
		record, ok := data[i].(model.Record)
		if !ok {
			log.Error().Msg("error cast record")
		}
		if !pkg.IsRecordExisted(record, *validRecords) {
			*validRecords = append(*validRecords, record)
		}
	}

	return nil
}

// Close closed connection
func (s *tickDBRepository) Close() error {
	return s.db.Close()
}

func parseCandleRecords(row *sql.Rows) []interface{} {
	var result []interface{}
	var record model.Record
	for row.Next() {
		err := row.Scan(&record.Type, &record.Kubun, &record.Hasshin, &record.CreateDay, &record.CreateTime, &record.StartIndex, &record.PathFolder)
		if err != nil {
			log.Error().Msg(err.Error())

			return nil
		}
		result = append(result, record)
	}

	return result
}
