/*
Package starttask implements logics start ECS task.
*/
package starttask

import (
	"context"
	"toiawase-start/configs"
	"toiawase-start/infrastructure/repository"
)

type service struct {
	config        *configs.Server
	ecsRepository repository.IECSRepository
}

// NewService constructor init service start ECS task
func NewService(cfg *configs.Server, ecs repository.IECSRepository) IStartTask {
	return &service{
		config:        cfg,
		ecsRepository: ecs,
	}
}

// StartService start service update task
func (s *service) StartService(ctx context.Context) error {
	return s.ecsRepository.UpdateTask(ctx)
}
