/*
Package api implements logics management of services.
*/
package api

import (
	"chikuseki-check/configs"
	"chikuseki-check/infrastructure"
	"chikuseki-check/infrastructure/repository"
	countdata "chikuseki-check/usecase/count_data"
	loadconfig "chikuseki-check/usecase/load_config"
	"context"
	"github.com/rs/zerolog/log"
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
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	if err := loadConfigService.LoadConfig(); err != nil {
		return err
	}

	log.Info().Msg("start chikuseki check process")
	countService := countdata.NewService(&s.cfg.TickSystem, tickDBRepository)

	return countService.CountData(ctx)
}

// Close closes server and related resources
func (s *Server) Close() {
	log.Info().Msg("release server resources")
	defer log.Info().Msg("completed release server resources")

	s.infra.Close()
}
