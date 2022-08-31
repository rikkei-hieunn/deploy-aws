/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"context"
	"update-status/configs"
	"update-status/infrastructure/aws/s3"
)

// Infra infrastructure management struct
type Infra struct {
	S3Handler s3.IS3Handler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler: s3Handler,
	}, nil
}
