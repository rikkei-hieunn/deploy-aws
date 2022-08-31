/*
Package api implements logics management of services.
*/
package api

import (
	"context"
	"send-command/configs"
	"send-command/infrastructure"
	"send-command/infrastructure/repository"
	"send-command/model"
	sendrequest "send-command/usecase/send_request"
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
	// Create AWS TickSystem repository
	s3Repository := repository.NewStorageRepository(s.infra)
	socketRepository := repository.NewSocketRepository(s.infra)

	var sender sendrequest.ISender
	if s.cfg.TickSystem.RequestType == model.RequestTypeAll {
		sender = sendrequest.NewSendAllService(&s.cfg.TickSystem, s3Repository, socketRepository)
	} else if s.cfg.TickSystem.RequestType == model.RequestTypeGroupLine {
		sender = sendrequest.NewSendLineService(&s.cfg.TickSystem, s3Repository, socketRepository)
	} else if s.cfg.TickSystem.RequestType == model.RequestTypeGroupID {
		sender = sendrequest.NewSendGroupService(&s.cfg.TickSystem, s3Repository, socketRepository)
	} else if s.cfg.TickSystem.RequestType == model.RequestTypeToiawase {
		sender = sendrequest.NewSendToiawaseService(&s.cfg.TickSystem, socketRepository)
	}

	if sender == nil {
		return nil
	}

	return sender.HandleRequest()
}
