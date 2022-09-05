/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"ec2-stop/configs"
	"ec2-stop/infrastructure"
	"ec2-stop/infrastructure/repository"
	stopinstance "ec2-stop/usecase/stop_instance"
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
func (s *Server) Start(ctx context.Context) error {
	// Create AWS TickSystem repository
	ec2Repository := repository.NewEC2Repository(s.infra)
	stopInstanceService := stopinstance.NewService(&s.cfg.TickSystem, ec2Repository)

	return stopInstanceService.Start(ctx)
}
