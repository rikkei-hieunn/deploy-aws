/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"chikuseki-check/configs"
	"chikuseki-check/infrastructure/aws/s3"
	"chikuseki-check/infrastructure/db/rdb"
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
