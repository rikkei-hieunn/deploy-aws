/*
Package handledel implements logics query receive.
*/
package handledel

import "context"

// IRequestHandler interface receiver
type IRequestHandler interface {
	Start(ctx context.Context,dbType string) error
}
