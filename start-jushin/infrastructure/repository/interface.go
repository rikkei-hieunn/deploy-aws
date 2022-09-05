/*
Package repository implements logics repository.
*/
package repository

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

//ISSMRepository provides all services about SSM
type ISSMRepository interface {
	ExecuteProgram(ctx context.Context)  (map[string]types.CommandInvocationStatus, error)
}
// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}
