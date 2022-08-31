package loadconfig

import (
	"context"
	"fmt"
	"os"
	"start-ecs/configs"
	"start-ecs/model"
	"strings"
)

type service struct {
	configs *configs.Server
}

//NewService constructor
func NewService(configs *configs.Server) ILoadConfig {
	return &service{
		configs: configs,
	}
}

//LoadConfig start load config
func (s service) LoadConfig(ctx context.Context, serviceName string) error {
	var envKeys, params []string
	var isEnoughParams bool
	runningType := os.Args[2]

	if serviceName == model.BP03 {
		if runningType == model.BP03FirstRunningType {
			envKeys = s.configs.BP03FirstRunningTypeEnvKeys
		}
		if runningType == model.BP03SecondRunningType {
			envKeys = s.configs.BP03SecondRunningTypeEnvKeys
		}
	} else if serviceName == model.BP05 {
		switch runningType {
		case model.BP05DemegetRunningType, model.BP05DemegetEvRunningType, model.BP05DemegetPtsRunningType, model.BP05DemegetPtsEvRunningType:
			envKeys = s.configs.BP05FirstRunningTypeEnvKeys
		case model.BP05DownloadAllRunningType, model.BP05DownloadAllEvoRunningType, model.BP05DownloadAllPtsRunningType, model.BP05DownloadAllEvPtsRunningType:
			envKeys = s.configs.BP05SecondRunningTypeEnvKeys
		case model.BP05RecreateRunningType, model.BP05RecreatePtsRunningType:
			envKeys = s.configs.BP05ThirdRunningTypeEnvKeys
		case model.BP05ReDownloadRunningType, model.BP05ReDownloadEvRunningType:
			envKeys = s.configs.BP05FourthRunningTypeEnvKeys
		}
	} else if serviceName == model.BP06 {
		envKeys = s.configs.BP06EnvKeyParams
	} else if serviceName == model.BP07 {
		if runningType == model.BP07FirstRunningType {
			envKeys = s.configs.BP07FirstRunningTypeEnvKeys
		}
		if runningType == model.BP07SecondRunningType {
			envKeys = s.configs.BP07SecondRunningTypeEnvKeys
		}
		if runningType == model.BP07ThirdRunningType {
			envKeys = s.configs.BP07ThirdRunningTypeEnvKeys
		}
	} else {
		return fmt.Errorf("can not found services to run : services name %s ", serviceName)
	}

	isEnoughParams = len(os.Args)-2 == len(envKeys)
	if !isEnoughParams {
		return fmt.Errorf("number of env argument and input params not match in process for : %s", serviceName)
	}

	params = os.Args[2:]
	cluster := strings.TrimSpace(os.Getenv(serviceName + model.UnderscoreCharacter + model.ClusterName))
	if cluster == model.EmptyString {
		return fmt.Errorf("cluster not found")
	}
	taskDefinition := strings.TrimSpace(os.Getenv(serviceName + model.UnderscoreCharacter + model.TaskDefinition))
	if taskDefinition == model.EmptyString {
		return fmt.Errorf("task definition not found")
	}
	containerName := strings.TrimSpace(os.Getenv(serviceName + model.UnderscoreCharacter + model.ContainerName))
	if containerName == model.EmptyString {
		return fmt.Errorf("container name not found")
	}
	subnets := strings.Split(os.Getenv(model.Subnets), model.CommaCharacter)
	if len(subnets) == 0 {
		return fmt.Errorf("subnets name not found")
	}
	securityGroups := strings.Split(os.Getenv(model.SecurityGroups), model.CommaCharacter)
	if len(securityGroups) == 0 {
		return fmt.Errorf("security groups name not found")
	}

	ecs := &configs.ECS{
		ClusterName:    cluster,
		TaskDefinition: taskDefinition,
		Subnets:        subnets,
		SecurityGroups: securityGroups,
		ContainerName:  containerName,
	}

	s.configs.EnvVarKeys = envKeys
	s.configs.EnvVarValues = params
	s.configs.ECS = ecs

	return nil
}
