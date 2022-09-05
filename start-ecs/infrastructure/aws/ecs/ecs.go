package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"start-ecs/configs"
	"start-ecs/model"
	"strings"
)

type ecsClient struct {
	config *configs.Server
	client *ecs.Client
}

// NewECSClient construct
func NewECSClient(ctx context.Context, cfg *configs.Server) IECSHandler {
	conf, _ := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.S3Region))
	client := ecs.NewFromConfig(conf)

	return &ecsClient{
		config: cfg,
		client: client,
	}
}

func (e *ecsClient) StartTask(ctx context.Context) (*string, error) {
	runTaskInput := e.setRunTaskInput()

	output, err := e.client.RunTask(ctx, runTaskInput)
	if output == nil {
		return nil, err
	}
	taskID := *output.Tasks[0].TaskArn
	lastIndex := strings.LastIndex(*output.Tasks[0].TaskArn, model.StrokeCharacter)
	taskID = taskID[lastIndex+1:]

	return &taskID, err
}

func (e *ecsClient) DescribeTask(ctx context.Context, taskID *string) (*string, error) {
	describeTaskInput := ecs.DescribeTasksInput{
		Tasks:   []string{*taskID},
		Cluster: &e.config.ClusterName,
	}
	tasksOutput, err := e.client.DescribeTasks(ctx, &describeTaskInput)
	if err != nil {
		return nil, err
	}

	return tasksOutput.Tasks[0].LastStatus, nil
}

func (e *ecsClient) StopTask(ctx context.Context, taskID *string) (*string, error) {
	stopTaskInput := ecs.StopTaskInput{
		Task:    taskID,
		Cluster: &e.config.ClusterName,
	}
	stopTaskOutput, err := e.client.StopTask(ctx, &stopTaskInput)
	if err != nil {
		return nil, err
	}

	return stopTaskOutput.Task.LastStatus, nil
}
func (e *ecsClient) setRunTaskInput() *ecs.RunTaskInput {
	var envKeyValues []types.KeyValuePair
	awsVpc := types.AwsVpcConfiguration{
		Subnets:        e.config.Subnets,
		SecurityGroups: e.config.SecurityGroups,
		AssignPublicIp: types.AssignPublicIpEnabled,
	}
	network := types.NetworkConfiguration{
		AwsvpcConfiguration: &awsVpc,
	}
	for i := 0; i < len(e.config.EnvVarKeys); i++ {
		keyValuePair := types.KeyValuePair{
			Name:  &e.config.EnvVarKeys[i],
			Value: &e.config.EnvVarValues[i],
		}
		envKeyValues = append(envKeyValues, keyValuePair)
	}
	containerOverride := types.ContainerOverride{
		Name:        &e.config.ContainerName,
		Environment: envKeyValues,
	}
	taskOverride := types.TaskOverride{
		ContainerOverrides: []types.ContainerOverride{containerOverride},
	}

	return &ecs.RunTaskInput{
		Cluster:              &e.config.ClusterName,
		TaskDefinition:       &e.config.TaskDefinition,
		NetworkConfiguration: &network,
		LaunchType:           types.LaunchTypeFargate,
		Overrides:            &taskOverride,
	}
}
