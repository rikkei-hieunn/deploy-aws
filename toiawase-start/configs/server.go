/*
Package configs contains configuration info.
*/
package configs

import (
	"fmt"
	"os"
	"strconv"
	"toiawase-start/model"
)

// Server application settings
type Server struct {
	ECS
}

// Init application configuration
func Init() (*Server, error) {
	ecs, err := initECSConfig()
	if err != nil {
		return nil, err
	}

	return &Server{
		ECS: *ecs,
	}, nil
}

func initECSConfig() (*ECS, error) {
	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return nil, fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	cluster := os.Getenv(model.ClusterName)
	if cluster == model.EmptyString {
		return nil, fmt.Errorf("invalid ECS cluster name")
	}

	service := os.Getenv(model.ServiceName)
	if service == model.EmptyString {
		return nil, fmt.Errorf("invalid ECS service name")
	}

	taskCountString := os.Getenv(model.TaskCount)
	if taskCountString == model.EmptyString {
		return nil, fmt.Errorf("invalid ECS task count")
	}

	taskCount, err := strconv.ParseInt(taskCountString, 10, 32)
	if err != nil {
		return nil, err
	}

	return &ECS{
		Region:    region,
		Cluster:   cluster,
		Service:   service,
		TaskCount: int32(taskCount),
	}, nil
}
