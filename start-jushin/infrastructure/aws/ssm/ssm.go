package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"start-jushin/configs"
	"start-jushin/model"
	"time"
)

type ssmClient struct {
	config *configs.Server
	client *ssm.Client
	waiter *ssm.CommandExecutedWaiter
}

// NewSSMClient constructor creates new ssm client in aws
func NewSSMClient(cfg *configs.Server) (ISSMHandler, error) {
	conf, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.S3Region))
	if err != nil {
		return nil, err
	}
	client := ssm.NewFromConfig(conf)

	return &ssmClient{
		config: cfg,
		client: client,
		waiter: ssm.NewCommandExecutedWaiter(client),
	}, nil
}

func (e *ssmClient) SendCommand(ctx context.Context) (map[string]types.CommandInvocationStatus, error) {
	parameterName := "AWS-RunShellScript"
	key := "InstanceIds"
	var commandID *string
	result := make(map[string]types.CommandInvocationStatus)
	commandInput := &ssm.SendCommandInput{
		DocumentName: &parameterName,
		Targets: []types.Target{
			{
				Key:    &key,
				Values: e.config.InstancesIDs,
			},
		},
		Parameters: map[string][]string{
			"commands": e.config.Commands,
		},
	}

	results, err := e.client.SendCommand(ctx, commandInput)
	if results == nil {
		return nil, err
	}

	//check status for each instance
	commandID = results.Command.CommandId
	commandInvokeInput := ssm.GetCommandInvocationInput{
		CommandId: commandID,
	}

	for i := range e.config.InstancesIDs {
		commandInvokeInput.InstanceId = &e.config.InstancesIDs[i]
		output, err := e.waiter.WaitForOutput(ctx, &commandInvokeInput, 15*time.Minute, e.setCommandExecutedWaiter)

		if output != nil && err == nil {
			result[e.config.InstancesIDs[i]] = output.Status

			continue
		}
		result[e.config.InstancesIDs[i]] = model.EmptyString
	}

	return result, err
}

func (e *ssmClient) setCommandExecutedWaiter(option *ssm.CommandExecutedWaiterOptions) {
	//TODO move config
	option.MinDelay = time.Duration(5000)
	option.MaxDelay = time.Duration(5000)
}
