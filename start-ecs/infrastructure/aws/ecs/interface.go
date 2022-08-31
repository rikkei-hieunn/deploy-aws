/*
Package ecs provides all about ecs
*/
package ecs

import "context"

//IECSHandler provides all services
type IECSHandler interface {
	StartTask(context.Context) (*string, error)
	DescribeTask(ctx context.Context, taskID *string) (*string, error)
	StopTask(ctx context.Context, taskID *string) (*string, error)
}
