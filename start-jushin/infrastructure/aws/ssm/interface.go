/*
Package ssm implements logics about ssm
*/
package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// ISSMHandler provides services for ssm handler
type ISSMHandler interface {
	SendCommand(ctx context.Context) (map[string]types.CommandInvocationStatus, error)
}
