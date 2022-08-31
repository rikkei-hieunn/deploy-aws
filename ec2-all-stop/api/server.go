/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"ec2-all-stop/configs"
	"ec2-all-stop/infrastructure"
	"ec2-all-stop/infrastructure/repository"
	loadconfig "ec2-all-stop/usecase/load_config"
	stopinstance "ec2-all-stop/usecase/stop_instance"
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
	s3Repository := repository.NewStorageRepository(s.infra)

	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	err := loadConfigService.LoadConfig(ctx)
	if err != nil {
		return err
	}
	stopInstanceService := stopinstance.NewService(&s.cfg.EC2, ec2Repository)

	return stopInstanceService.Start(ctx)
}
