/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"show-status/configs"
	"show-status/infrastructure"
	"show-status/infrastructure/repository"
	loaddata "show-status/usecase/load_data"
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

	// Load data from database status file
	loadDataService := loaddata.NewService(s.cfg, s3Repository)

	return loadDataService.LoadDatabaseStatus()
}

// Close closes server and related resources
func (s *Server) Close() {
	log.Info().Msg("release server resources")
	defer log.Info().Msg("completed release server resources")
}
