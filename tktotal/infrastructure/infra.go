// Package infrastructure for application
package infrastructure

import (
	"context"
	"tktotal/configs"
	"tktotal/infrastructure/storage/s3"
)

//Infra structure all internal services
type Infra struct {
	S3Handler s3.IS3Handler
}

//Init connect internal service
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	s3Handler, err := s3.NewS3Client(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler: s3Handler,
	}, nil
}
