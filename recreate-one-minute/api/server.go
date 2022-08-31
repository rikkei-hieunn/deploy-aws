// Package api for server
package api

import (
	"context"
	"fmt"
	"recreate-one-minute/configs"
	"recreate-one-minute/infrastructure"
	"recreate-one-minute/infrastructure/repository"
	handlerecreate "recreate-one-minute/usecase/handle_recreate"
	loadconfig "recreate-one-minute/usecase/load_config"
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

// Start func start server instance
func (s Server) Start(ctx context.Context) error {
	var err error
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)

	//Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	err = loadConfigService.LoadConfig(ctx)
	if err != nil {
		return fmt.Errorf("service load config error %w : ", err)
	}

	//init database
	err = s.infra.InitDatabase(s.cfg)
	if err != nil {
		return err
	}

	// Create TICK DB repository
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	service := handlerecreate.NewService(s.cfg, tickDBRepository)

	return service.Start(ctx)
}
