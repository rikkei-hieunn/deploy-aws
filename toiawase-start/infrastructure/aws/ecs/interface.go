package ecs

import (
	"context"
)

// IECSHandler method about ECS handler
type IECSHandler interface {
	UpdateTask(context.Context) error
}
