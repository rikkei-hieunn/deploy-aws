/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"create-table/configs"
	"create-table/infrastructure"
	"create-table/infrastructure/repository"
	"create-table/model"
	"create-table/pkg/logger"
	createtable "create-table/usecase/create_table"
	loadconfig "create-table/usecase/load_config"
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
func (s *Server) Start(ctx context.Context) error {
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	if err := loadConfigService.LoadConfig(); err != nil {
		return err
	}

	log.Info().Msg("start create table process")
	wg := &sync.WaitGroup{}
	if s.cfg.TickSystem.CreateType == model.CreateTypeKei1 || s.cfg.TickSystem.CreateType == model.CreateTypeBoth {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createTableService := createtable.NewService(&s.cfg.TickSystem, tickDBRepository)
			errs := createTableService.CreateTables(ctx, model.TargetCreateTableTheFirstKei)
			if errs != nil {
				// TODO send to Senjiu system
				logger.ShowLog(errs)
			}
		}()
	}
	if s.cfg.TickSystem.CreateType == model.CreateTypeKei2 || s.cfg.TickSystem.CreateType == model.CreateTypeBoth {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createTableService := createtable.NewService(&s.cfg.TickSystem, tickDBRepository)
			errs := createTableService.CreateTables(ctx, model.TargetCreateTableTheSecondKei)
			if errs != nil {
				// TODO send to Senjiu system
				logger.ShowLog(errs)
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
