/*
Package infrastructure init config SortingDB, DateTime, Files.
*/
package infrastructure

import (
	"insert-calendar-infos/configs"
	"insert-calendar-infos/infrastructure/db/rdb"
	"insert-calendar-infos/infrastructure/filebus"
	"insert-calendar-infos/infrastructure/s3"
)

// Infra infrastructure management struct
type Infra struct {
	FilebusHandler filebus.IFilebusHandler
	TickDB         rdb.ITickDBHandler
	S3Handler      s3.IS3Handler
}

// Init initializes resources
func Init(cfg *configs.Server) (*Infra, error) {
	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	filebusHandler := filebus.NewFilebusHandler(&cfg.TickFileBus)

	return &Infra{
		S3Handler:      s3Handler,
		FilebusHandler: filebusHandler,
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
func (i *Infra) Close() {
	if i.TickDB != nil {
		_ = i.TickDB.Close()
	}
}
