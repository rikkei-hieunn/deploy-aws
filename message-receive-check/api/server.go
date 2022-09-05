package api

import (
	"context"
	"fmt"
	"message-receive-check/configs"
	"message-receive-check/infrastructure"
	"message-receive-check/infrastructure/repository"
	checkmessages "message-receive-check/usecase/check_messages"
	loadconfig "message-receive-check/usecase/load_config"
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
func (s Server) Start(ctx context.Context) {
	// Create AWS S3 repository
	s3Repository := repository.NewStorageRepository(s.infra)

	// Load configuration file
	loadConfigService := loadconfig.NewService(s.cfg, s3Repository)
	err := loadConfigService.LoadConfig(ctx)
	if err != nil {
		// TODO handle log err
		fmt.Println(err.Error())
	}

	checkMessagesService := checkmessages.NewService(s.cfg, s3Repository)
	errs := checkMessagesService.Start(ctx)
	if errs != nil {
		// TODO handle log err
		for index := range errs {
			fmt.Println(errs[index].Error())
		}
	}
}
