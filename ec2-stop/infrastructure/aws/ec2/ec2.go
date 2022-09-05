package ec2

import (
	"context"
	"ec2-stop/configs"
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
func NewEC2Handler(ctx context.Context, cfg *configs.TickSystem) (IEC2Handler, error) {
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(conf)

	return &ec2Client{
		client: client,
	}, nil
}

// StopInstance stop an instance with id
func (e *ec2Client) StopInstance(ctx context.Context, instanceID string) error {
	var apiErr smithygo.APIError
	dryRun := true
	input := &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
		DryRun:      &dryRun,
	}
	_, err := e.client.StopInstances(ctx, input)

	dryRun = false
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = &dryRun
		_, err = e.client.StopInstances(ctx, input)
	}

	return err
}

// StopInstances start many instances
func (e *ec2Client) StopInstances(ctx context.Context, instanceIDs []string) error {
	var apiErr smithygo.APIError
	dryRun := true
	input := &ec2.StopInstancesInput{
		InstanceIds: instanceIDs,
		DryRun:      &dryRun,
	}
	_, err := e.client.StopInstances(ctx, input)

	dryRun = false
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = &dryRun
		_, err = e.client.StopInstances(ctx, input)
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
