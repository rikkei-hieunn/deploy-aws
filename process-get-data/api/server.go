/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"process-get-data/configs"
	"process-get-data/infrastructure"
	"process-get-data/infrastructure/repository"
	"process-get-data/model"
	insertdata "process-get-data/usecase/insert_data"
	loadconfig "process-get-data/usecase/load_config"
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
func (s *Server) Start(ctx context.Context) error {
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	if err := loadConfigService.LoadConfig(); err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	if s.cfg.TickSystem.ProcessGetData == model.GetDataTypeKei1 || s.cfg.TickSystem.ProcessGetData == model.GetDataTypeBoth {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createJobsService := insertdata.NewService(&s.cfg.TickSystem, tickDBRepository)
			err := createJobsService.InsertData(ctx, model.QuoteCodesDefinitionTheFirstKei, model.GetDataTypeKei1)
			if err != nil {
				// TODO send to Senjiu system
			}
		}()
	}
	if s.cfg.TickSystem.ProcessGetData == model.GetDataTypeKei2 || s.cfg.TickSystem.ProcessGetData == model.GetDataTypeBoth {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createJobsService := insertdata.NewService(&s.cfg.TickSystem, tickDBRepository)
			err := createJobsService.InsertData(ctx, model.QuoteCodesDefinitionTheSecondKei, model.GetDataTypeKei2)
			if err != nil {
				// TODO send to Senjiu system
			}
		}()
	}
	wg.Wait()

	return nil
}

// Close closes server and related resources
func (s *Server) Close() {
	log.Info().Msg("release server resources")
	defer log.Info().Msg("completed release server resources")

	s.infra.Close()
}
