package starttask

import (
	"context"
)

// IStartTask methods start ECS task
type IStartTask interface {
	StartService(ctx context.Context) error
}
