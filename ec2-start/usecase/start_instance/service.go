package startinstance

import (
	"context"
	"ec2-start/configs"
	"ec2-start/infrastructure/repository"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/rs/zerolog/log"
)

type service struct {
	cfg           *configs.TickSystem
	ec2Repository repository.IEC2Repository
}

// NewService constructor init service start TickSystem instances
func NewService(config *configs.TickSystem, ec2Repository repository.IEC2Repository) IStartInstance {
	return &service{
		cfg:           config,
		ec2Repository: ec2Repository,
	}
}

// Start service
func (s *service) Start(ctx context.Context) error {
	statuses, err := s.ec2Repository.GetStatus(ctx, s.cfg.InstanceIds)
	if err != nil {
		return err
	}

	var validInstanceIDs []string
	for instanceID, status := range statuses {
		if status != types.InstanceStateNameStopped {
			log.Error().Msgf("cannot start instance, instance's id: %s, status: %s", instanceID, status)

			continue
		}

		validInstanceIDs = append(validInstanceIDs, instanceID)
	}

	if len(validInstanceIDs) == 0 {
		return fmt.Errorf("no instances could be started")
	}

	err = s.ec2Repository.StartInstances(ctx, validInstanceIDs)
	if err != nil {
		return err
	}

	return nil
}
