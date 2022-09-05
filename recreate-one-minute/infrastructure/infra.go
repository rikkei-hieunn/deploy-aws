/*
Package infrastructure init config SortingDB, DateTime, Files.
*/
package infrastructure

import (
	"github.com/rs/zerolog/log"
	"recreate-one-minute/configs"
	"recreate-one-minute/infrastructure/db/rdb"
	"recreate-one-minute/infrastructure/storage/s3"
)

// Infra infrastructure management struct
type Infra struct {
	TickDB    rdb.ITickDBHandler
	S3Handler s3.IS3Handler
}

// Init initializes resources
func Init(cfg *configs.Server) (*Infra, error) {
	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler: s3Handler,
	}, nil
}

// InitDatabase init database connection
func (r *Infra) InitDatabase(cfg *configs.Server) error {
	tickDB, err := rdb.NewTickDBHandler(&cfg.TickDB)
	if err != nil {
		return err
	}
	r.TickDB = tickDB

	return nil
}

// Close closes resources gracefully
func (r *Infra) Close() {
	if r.TickDB != nil {
		if err := r.TickDB.Close(); err != nil {
			log.Warn().Msg(err.Error())
		}
	}
}
