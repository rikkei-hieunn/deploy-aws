/*
Package infrastructure init config SortingDB, DateTime, Files.
*/
package infrastructure

import (
	"context"
	"data-del/configs"
	"data-del/infrastructure/db/rdb"
	"data-del/infrastructure/storage/s3"
	"github.com/rs/zerolog/log"
)

// Infra infrastructure management struct
type Infra struct {
	TickDB    rdb.ITickDBHandler
	S3Handler s3.IS3Handler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	s3Handler, err := s3.NewS3Client(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler: s3Handler,
	}, nil
}

// Close closes resources gracefully
func (r *Infra) Close() {
	if r.TickDB != nil {
		if err := r.TickDB.Close(); err != nil {
			log.Warn().Msg(err.Error())
		}
	}
}

//InitDatabase init database
func (r *Infra) InitDatabase(tickDBConfigs *configs.TickDB) error {
	tickDB, err := rdb.NewTickDBHandler(tickDBConfigs)
	if err != nil {
		return err
	}
	r.TickDB = tickDB

	return nil
}
