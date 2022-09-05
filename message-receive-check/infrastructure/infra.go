package infrastructure

import (
	"context"
	"message-receive-check/configs"
	"message-receive-check/infrastructure/aws/s3"
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
