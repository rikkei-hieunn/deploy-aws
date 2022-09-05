/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"context"
	"send-command/configs"
	"send-command/infrastructure/aws/s3"
	"send-command/infrastructure/socket"
)

// Infra infrastructure management struct
type Infra struct {
	S3Handler     s3.IS3Handler
	SocketHandler socket.ISocketHandler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	socketHandler := socket.NewSocketHandler()
	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{
		S3Handler:     s3Handler,
		SocketHandler: socketHandler,
	}, nil
}
