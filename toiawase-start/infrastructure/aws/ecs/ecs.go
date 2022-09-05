/*
Package ecs implements logics about ecs.
*/
package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"toiawase-start/configs"
)

type ecsHandler struct {
	config *configs.ECS
	client *ecs.Client
}

// NewECSHandler constructor init handler
func NewECSHandler(ctx context.Context, cfg *configs.ECS) (IECSHandler, error) {
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
	if err != nil {
		return nil, err
	}
	client := ecs.NewFromConfig(conf)

	return &ecsHandler{
		config: cfg,
		client: client,
	}, nil
}

// UpdateTask update task ecs
func (e *ecsHandler) UpdateTask(ctx context.Context) error {
	updateServiceInput := ecs.UpdateServiceInput{
		Cluster:      &e.config.Cluster,
		Service:      &e.config.Service,
		DesiredCount: &e.config.TaskCount,
	}
	_, err := e.client.UpdateService(ctx, &updateServiceInput)

	return err
}
