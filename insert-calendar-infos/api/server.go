/*
Package api init server infrastructure, config.
*/
package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/infrastructure"
	"insert-calendar-infos/infrastructure/repository"
	"insert-calendar-infos/model"
	handlerequest "insert-calendar-infos/usecase/handle_request"
	loadconfig "insert-calendar-infos/usecase/load_config"
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

	// Create filebus repository
	filebusRepository := repository.NewFilebusRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	err := loadConfigService.LoadConfig()
	if err != nil {
		return err
	}

	// init database connection
	err = s.infra.InitDatabase(s.cfg)
	if err != nil {
		return err
	}

	// Create TICK DB repository
	tickDBRepository := repository.NewTickDBRepository(s.infra)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Start handle request for the first kei
	requestHandlerServiceFirstKei := handlerequest.NewService(s.cfg, tickDBRepository, filebusRepository, s3Repository)
	go func() {
		defer wg.Done()
		err = requestHandlerServiceFirstKei.Start(ctx, model.TheFirstKei)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}()

	// Start handle request for the second kei
	requestHandlerServiceSecondKei := handlerequest.NewService(s.cfg, tickDBRepository, filebusRepository, s3Repository)
	go func() {
		defer wg.Done()
		err = requestHandlerServiceSecondKei.Start(ctx, model.TheSecondKei)
		if err != nil {
			log.Error().Msg(err.Error())
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
