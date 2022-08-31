package check_messages

import (
	"context"
)

// ICheckMessages interface define methods of service check message
type ICheckMessages interface {
	Start(ctx context.Context) []error
}

// IWorker interface define methods of worker
type IWorker interface {
	Start(ctx context.Context, s string) error
}
