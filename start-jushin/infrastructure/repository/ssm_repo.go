package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"start-jushin/infrastructure/aws/ssm"
)

type ssmRepository struct {
	ssmHandler ssm.ISSMHandler
}

// NewSSMClient constructor creates new ssm client instance
func NewSSMClient(ssmHandler ssm.ISSMHandler) ISSMRepository {
	return &ssmRepository{ssmHandler: ssmHandler}
}

func (e *ssmRepository) ExecuteProgram(ctx context.Context) (map[string]types.CommandInvocationStatus, error) {
	return e.ssmHandler.SendCommand(ctx)
}
