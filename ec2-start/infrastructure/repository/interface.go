/*
Package repository implements logics repository.
*/
package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// IEC2Repository list of actions EC2 repository
type IEC2Repository interface {
	StartInstance(ctx context.Context, id string) error
	StartInstances(ctx context.Context, ids []string) error
	GetStatus(ctx context.Context, instanceIDs []string) (map[string]types.InstanceStateName, error)
}
