/*
Package api start server to start all use case
*/
package api

import (
	"context"
	"fmt"
	"tktotal/configs"
	"tktotal/infrastructure"
	"tktotal/infrastructure/repository"
	"tktotal/usecase/load_config"
	summarizelog "tktotal/usecase/summarize_log"
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

//Start server
func (s Server) Start(ctx context.Context, date string) error {
	var err error
	s3Repository := repository.NewStorageRepository(s.infra)
	loadConfigSvc := loadconfig.NewService(s.cfg, s3Repository)

	err = loadConfigSvc.LoadConfig(ctx,date)
	if err != nil {
		return fmt.Errorf("load config fail %w ", err)
	}

	summarizeLogService := summarizelog.NewService(s.cfg, s3Repository)
	err = summarizeLogService.Start(ctx)
	if err != nil {
		return fmt.Errorf("summerize log fail %w ", err)
	}

	return err
}
