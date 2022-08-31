/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"toiawase-start/configs"
	"toiawase-start/infrastructure"
	"toiawase-start/infrastructure/repository"
	starttask "toiawase-start/usecase/start_task"
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

// Start starts server and related resources
func (s Server) Start(ctx context.Context) error {
	ecsRepository := repository.NewECSRepository(s.infra)
	startTaskService := starttask.NewService(s.cfg, ecsRepository)

	return startTaskService.StartService(ctx)
}
