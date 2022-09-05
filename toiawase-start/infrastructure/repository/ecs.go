package repository

import (
	"context"
	"toiawase-start/infrastructure"
	ecs2 "toiawase-start/infrastructure/aws/ecs"
)

type ecsRepository struct {
	ecsHandler ecs2.IECSHandler
}

// NewECSRepository constructor init ECS repository
func NewECSRepository(infra *infrastructure.Infra) IECSRepository {
	return &ecsRepository{ecsHandler: infra.ECSHandler}
}

// UpdateTask update task ecs
func (e *ecsRepository) UpdateTask(ctx context.Context) error {
	return e.ecsHandler.UpdateTask(ctx)
}
