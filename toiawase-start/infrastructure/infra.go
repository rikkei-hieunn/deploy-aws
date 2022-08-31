/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"context"
	"toiawase-start/configs"
	"toiawase-start/infrastructure/aws/ecs"
)

// Infra infrastructure management struct
type Infra struct {
	ECSHandler ecs.IECSHandler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	ecsHandler, err := ecs.NewECSHandler(ctx, &cfg.ECS)
	if err != nil {
		return nil, err
	}

	return &Infra{ECSHandler: ecsHandler}, nil
}
