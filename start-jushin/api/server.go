// Package api server
package api

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"start-jushin/configs"
	"start-jushin/infrastructure"
	"start-jushin/infrastructure/repository"
	"start-jushin/usecase/load_config"
	"start-jushin/usecase/start_instance"
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

//Start func start server
func (s Server) Start(ctx context.Context, startType, keiType, groupID, dataType, groupLine string) error {
	var err error

	s3Repository, err := repository.NewStorageRepository(s.infra)
	if err != nil {
		return fmt.Errorf("error while starting instance: %w ", err)
	}

	loadConfigSvc := loadconfig.NewService(s.cfg, s3Repository)

	err = loadConfigSvc.LoadConfig(startType, keiType, groupID, dataType, groupLine)
	if err != nil {
		return fmt.Errorf("error while starting instance: %w ", err)
	}

	ssmRepository := repository.NewSSMClient(s.infra.SSMHandler)
	executeProgramService := startinstance.NewService(ssmRepository)
	err = executeProgramService.ExecuteProgram(ctx)
	if err != nil {
		return err
	}
	log.Info().Msg("Start task success...")

	return nil
}
