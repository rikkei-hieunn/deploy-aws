/*
Package ec2 implements logics EC2 instances.
*/
package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// IEC2Handler list of actions with EC2
type IEC2Handler interface {
	StopInstance(ctx context.Context, instanceID string) error
	StopInstances(ctx context.Context, instanceIDs []string) error
	GetStatus(ctx context.Context, instanceIDs []string) (map[string]types.InstanceStateName, error)
}
