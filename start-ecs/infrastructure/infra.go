/*
Package infrastructure all logic for internal service
*/
package infrastructure

import (
	"context"
	"start-ecs/configs"
	"start-ecs/infrastructure/aws/ecs"
	"start-ecs/infrastructure/aws/s3"
)

//Infra construct all internal services
type Infra struct {
	ECSHandler ecs.IECSHandler
	S3Handler  s3.IS3Handler
}

//Init start create infra instance and connect internal services
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	ecsHandler := ecs.NewECSClient(ctx, cfg)

	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)

	if err != nil {
		return nil, err
	}

	return &Infra{ECSHandler: ecsHandler, S3Handler: s3Handler}, nil
}
