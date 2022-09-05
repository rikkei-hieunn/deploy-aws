package repository

import (
	"context"
	"start-ecs/infrastructure"
	"start-ecs/infrastructure/aws/ecs"
)

type ecsRepository struct {
	ecsHandler ecs.IECSHandler
}

//NewECSRepository repository constructor
func NewECSRepository(infra *infrastructure.Infra) IECSRepository {
	return &ecsRepository{ecsHandler: infra.ECSHandler}
}

//StartService start task service
func (e *ecsRepository) StartService(ctx context.Context) (*string, error) {
	return e.ecsHandler.StartTask(ctx)
}

//GetTaskStatus get status service
func (e *ecsRepository) GetTaskStatus(ctx context.Context, taskID *string) (*string, error) {
	return e.ecsHandler.DescribeTask(ctx, taskID)
}

//StopService stop task services
func (e *ecsRepository) StopService(ctx context.Context, taskID *string) (*string, error) {
	return e.ecsHandler.StopTask(ctx, taskID)
}
