/*
Package starttask all logic for services to run ecs
*/
package starttask

import (
	"context"
	"fmt"
	"start-ecs/configs"
	"start-ecs/infrastructure/repository"
	"start-ecs/model"
	"time"
)

type service struct {
	config        *configs.Server
	ecsRepository repository.IECSRepository
}

//NewService constructor
func NewService(cfg *configs.Server, ecs repository.IECSRepository) IStartTask {
	return &service{
		config:        cfg,
		ecsRepository: ecs,
	}
}

//Start run ecs service
func (s *service) Start(ctx context.Context) error {
	for retryTime := 0; retryTime < s.config.RetryWaitTime; retryTime++ {
		var taskStatus string
		taskID, err := s.ecsRepository.StartService(ctx)
		if err != nil {
			return err
		}
		for statusCount := 0; statusCount < s.config.RetryWaitTime; statusCount++ {
			time.Sleep(time.Duration(s.config.MaxWaitTime) * time.Millisecond)
			lastStatus, err := s.ecsRepository.GetTaskStatus(ctx, taskID)
			if err != nil {
				return err
			}
			if *lastStatus == model.TaskStatusRunning {
				return nil
			}
			if *lastStatus != taskStatus && taskStatus != model.EmptyString {
				statusCount = -1
				taskStatus = *lastStatus

				continue
			}
			if statusCount == s.config.MaxCountTime {
				lastStatus, err = s.ecsRepository.StopService(ctx, taskID)
				if err != nil {
					return err
				}
			}
		}
	}

	return fmt.Errorf("failed to start task")
}
