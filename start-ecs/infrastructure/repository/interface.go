/*
Package repository all logic for services
*/
package repository

import (
	"bytes"
	"context"
)

//IECSRepository provides all services about ECS
type IECSRepository interface {
	StartService(context.Context) (*string, error)
	GetTaskStatus(context.Context, *string) (*string, error)
	StopService(ctx context.Context, taskID *string) (*string, error)
}

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}
