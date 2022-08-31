/*
Package repository all logic for services
*/
package repository

import "context"

//IECSRepository provides all services about ECS
type IECSRepository interface {
	StartService(context.Context) (*string, error)
	GetTaskStatus(context.Context, *string) (*string, error)
	StopService(ctx context.Context, taskID *string) (*string, error)
}
