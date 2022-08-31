/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"update-status/configs"
	"update-status/infrastructure"
	"update-status/infrastructure/repository"
	"update-status/model"
	loadconfig "update-status/usecase/load_config"
	updatestatus "update-status/usecase/update_status"
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

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	if err := loadConfigService.LoadConfig(); err != nil {
		return err
	}

	var updater updatestatus.IUpdater
	if s.cfg.TickSystem.RequestType == model.UpdateStatusTypeQuoteCode {
		updater = updatestatus.NewUpdateTypeQuoteCodeService(&s.cfg.TickSystem, s3Repository)
	} else if s.cfg.TickSystem.RequestType == model.UpdateStatusTypeDBName {
		updater = updatestatus.NewUpdateTypeDBNameService(&s.cfg.TickSystem, s3Repository)
	} else if s.cfg.TickSystem.RequestType == model.UpdateStatusTypeGroupID {
		updater = updatestatus.NewUpdateTypeGroupIDService(&s.cfg.TickSystem, s3Repository)
	}

	if updater == nil {
		return nil
	}
	log.Info().Msg("start update process")
	defer log.Info().Msg("update is over")

	return updater.StartUpdateStatus()
}

// Close closes server and related resources
func (s *Server) Close() {
	log.Info().Msg("release server resources")
	defer log.Info().Msg("completed release server resources")
}
