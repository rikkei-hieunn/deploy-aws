/*
Package handlerecreate implements logics about receive message.
*/
package handlerecreate

import (
"context"
)

// IRequestHandler provides handle_request service interfaces
type IRequestHandler interface {
	Start(ctx context.Context) error
}