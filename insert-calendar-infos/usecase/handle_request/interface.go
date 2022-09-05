/*
Package handlerequest implements logics query receive.
*/
package handlerequest

import "context"

// IRequestHandler interface receiver
type IRequestHandler interface {
	Start(ctx context.Context, kei string) error
}
