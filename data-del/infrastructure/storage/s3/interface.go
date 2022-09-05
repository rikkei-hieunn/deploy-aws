/*
Package s3 implements connection se and contains config s3.
*/
package s3

import (
	"bytes"
	"context"
)

// IS3Handler interface
type IS3Handler interface {
	GetObject(ctx context.Context, path string) (*bytes.Buffer, error)
	CheckObjectExists(ctx context.Context, prefix string) (bool, error)
}
