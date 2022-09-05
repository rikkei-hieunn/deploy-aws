/*
Package stopinstance implements logics start EC2 instances.
*/
package stopinstance

import (
	"context"
)

// IStartInstance list of actions service
type IStartInstance interface {
	Start(ctx context.Context) error
}
