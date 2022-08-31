package repository

import (
	"os"
	"send-command/infrastructure"
	"send-command/infrastructure/socket"
	"send-command/model"
)

// socketRepository Structure of repository socket
type socketRepository struct {
	socketHandler socket.ISocketHandler
}

// NewSocketRepository Initialize a Repository socket
func NewSocketRepository(infra *infrastructure.Infra) ISocketRepository {
	return &socketRepository{
		socketHandler: infra.SocketHandler,
	}
}

// InitWriter init connection
func (s *socketRepository) InitWriter(machineName string, port int) error {
	host := os.Getenv(machineName + model.IPSuffix)

	return s.socketHandler.InitConnection(host, port)
}

// SendCommand socket send data
func (s *socketRepository) SendCommand(command string) error {
	return s.socketHandler.Send(command)
}
