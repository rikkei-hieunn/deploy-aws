/*
Package api start all use case services
*/
package api

import (
	"context"
	"data-del/configs"
	"data-del/infrastructure"
	"data-del/infrastructure/repository"
	"data-del/model"
	handledel "data-del/usecase/handle_delete"
	loadconfig "data-del/usecase/load_config"
	"github.com/rs/zerolog/log"
	"sync"
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
func (s *Server) Start(ctx context.Context, keiNumber string) error {
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	if err := loadConfigService.LoadConfig(ctx, keiNumber); err != nil {
		return err
	}
	//Init tick DB after load a list endpoint
	err := s.infra.InitDatabase(&s.cfg.TickDB)
	if err != nil {
		return err
	}
	// Create TICK DB repository
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	deleteService := handledel.NewService(&s.cfg.TickSystem, tickDBRepository, s3Repository)

	// Start receive request
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = deleteService.Start(ctx, model.TypeTick)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		err := deleteService.Start(ctx, model.TypeKehai)
		if err != nil {
			return
		}
	}()
	wg.Wait()

	return nil
}

// Close closes server and related resources
func (s *Server) Close() {
	log.Info().Msg("release server resources")
	defer log.Info().Msg("completed release server resources")

	s.infra.Close()
}
