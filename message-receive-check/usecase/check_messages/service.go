package check_messages

import (
	"context"
	"message-receive-check/configs"
	"message-receive-check/infrastructure/repository"
)

type service struct {
	cfg          *configs.Server
	s3Repository repository.IS3Repository
}

// NewService constructor create a new service
func NewService(config *configs.Server, s3Repo repository.IS3Repository) ICheckMessages {
	return &service{
		cfg:          config,
		s3Repository: s3Repo,
	}
}

// Start process check log file
func (s *service) Start(ctx context.Context) []error {
	var errs []error
	for index := range s.cfg.ProcessNames {
		worker := NewWorker(s.cfg, s.s3Repository)
		err := worker.Start(ctx, s.cfg.ProcessNames[index])
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
