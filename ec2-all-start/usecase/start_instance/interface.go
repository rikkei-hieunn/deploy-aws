/*
Package startinstance implements logics start EC2 instances.
*/
package startinstance

import (
	"context"
)

// IStartInstance list of actions service
type IStartInstance interface {
	Start(ctx context.Context) error
}
