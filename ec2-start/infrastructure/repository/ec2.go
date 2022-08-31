package repository

import (
	"context"
	"ec2-start/infrastructure"
	ec2 "ec2-start/infrastructure/aws/ec2"
	"ec2-start/model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// EC2Repository structure of EC2 repository
type EC2Repository struct {
	ec2Handler ec2.IEC2Handler
}

// NewEC2Repository constructor
func NewEC2Repository(infra *infrastructure.Infra) IEC2Repository {
	return &EC2Repository{
		ec2Handler: infra.EC2Handler,
	}
}

// StartInstance start an instance
func (e *EC2Repository) StartInstance(ctx context.Context, id string) error {
	if id == model.EmptyString {
		return fmt.Errorf("instance's id not found")
	}

	return e.ec2Handler.StartInstance(ctx, id)
}

// StartInstances start many instances
func (e *EC2Repository) StartInstances(ctx context.Context, ids []string) error {
	var validIDs []string
	for index := range ids {
		if ids[index] == model.EmptyString {
			continue
		}
		validIDs = append(validIDs, ids[index])
	}
	if validIDs == nil {
		return fmt.Errorf("instance's id not found")
	}

	return e.ec2Handler.StartInstances(ctx, validIDs)
}

// GetStatus get status of list instances
func (e *EC2Repository) GetStatus(ctx context.Context, instanceIds []string) (map[string]types.InstanceStateName, error) {
	return e.ec2Handler.GetStatus(ctx, instanceIds)
}
