/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"process-get-data/configs"
	"process-get-data/infrastructure/db/rdb"
	"process-get-data/infrastructure/s3"
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

	tickDB, err := rdb.NewTickDBHandler(&cfg.TickDB)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler: s3Handler,
		TickDB:    tickDB,
	}, nil
}

// Close closes resources gracefully
func (i *Infra) Close() {
	if i.TickDB != nil {
		_ = i.TickDB.Close()
	}
}
