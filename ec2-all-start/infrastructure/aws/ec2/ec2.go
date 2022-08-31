package ec2

import (
	"context"
	"ec2-all-start/configs"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithygo "github.com/aws/smithy-go"
)

type ec2Client struct {
	client *ec2.Client
}

// NewEC2Handler constructor init new EC2 handler
func NewEC2Handler(ctx context.Context, cfg *configs.Server) (IEC2Handler, error) {
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.TickSystem.S3Region))
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(conf)

	return &ec2Client{
		client: client,
	}, nil
}

// StartInstance start an instance with id
func (e *ec2Client) StartInstance(ctx context.Context, instanceID string) error {
	var apiErr smithygo.APIError
	dryRun := true
	input := &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
		DryRun:      &dryRun,
	}
	_, err := e.client.StartInstances(ctx, input)

	dryRun = false
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = &dryRun
		_, err = e.client.StartInstances(ctx, input)
	}

	return err
}

// StartInstances start many instances
func (e *ec2Client) StartInstances(ctx context.Context, instanceIDs []string) error {
	var apiErr smithygo.APIError
	dryRun := true
	input := &ec2.StartInstancesInput{
		InstanceIds: instanceIDs,
		DryRun:      &dryRun,
	}
	_, err := e.client.StartInstances(ctx, input)

	dryRun = false
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = &dryRun
		_, err = e.client.StartInstances(ctx, input)
	}

	return err
}

// GetStatus get status of instances
func (e *ec2Client) GetStatus(ctx context.Context, instanceIds []string) (map[string]types.InstanceStateName, error) {
	instanceStates := make(map[string]types.InstanceStateName)
	describeInput := ec2.DescribeInstancesInput{InstanceIds: instanceIds}
	output, err := e.client.DescribeInstances(ctx, &describeInput)
	if err != nil {
		return nil, err
	}

	for _, reservation := range output.Reservations {
		if len(reservation.Instances) == 0 {
			continue
		}
		instanceStates[*reservation.Instances[0].InstanceId] = reservation.Instances[0].State.Name
	}

	return instanceStates, err
}
