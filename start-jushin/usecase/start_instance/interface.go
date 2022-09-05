package startinstance

import (
	"context"
)

// IExecuteProgram provides services for start instance
type IExecuteProgram interface {
	ExecuteProgram(ctx context.Context) error
}
