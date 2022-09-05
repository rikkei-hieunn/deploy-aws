/*
Package api implements logic for server
*/
package api

import (
	"context"
	"fmt"
	"os"
	"start-ecs/configs"
	"start-ecs/infrastructure"
	"start-ecs/infrastructure/repository"
	loadconfig "start-ecs/usecase/load_config"
	starttask "start-ecs/usecase/start_task"
)

// Server instance
type Server struct {
	infra *infrastructure.Infra
	cfg   *configs.Server
}

// New create new server instance
func New(infra *infrastructure.Infra, cfg *configs.Server) *Server {
	return &Server{
		infra: infra,
		cfg:   cfg,
	}
}

//Start api server
func (s *Server) Start(ctx context.Context) error {
	var serviceName string
	var err error

	s3Repository, err := repository.NewStorageRepository(s.infra)
	if err != nil {
		return fmt.Errorf("error while starting instance: %w ", err)
	}

	loadConfigSvc := loadconfig.NewService(s.cfg, s3Repository)
	if serviceName = os.Args[1]; len(serviceName) == 0 {
		return fmt.Errorf("invalid first argument")
	}
	if err := loadConfigSvc.LoadConfig(ctx, serviceName); err != nil {
		return fmt.Errorf("load config fail %w ", err)
	}

	ecsRepo := repository.NewECSRepository(s.infra)

	runTaskSvc := starttask.NewService(s.cfg, ecsRepo)

	if err = runTaskSvc.Start(ctx); err != nil {
		return fmt.Errorf("run task error %w", err)
	}

	return nil
}
