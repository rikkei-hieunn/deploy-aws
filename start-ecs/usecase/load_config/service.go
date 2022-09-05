package loadconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"start-ecs/configs"
	"start-ecs/infrastructure/repository"
	"start-ecs/model"
	"strings"
)

type Service struct {
	Configs      *configs.Server
	S3Repository repository.IS3Repository
}

//NewService constructor
func NewService(configs *configs.Server, s3Repo repository.IS3Repository) ILoadConfig {
	return &Service{
		Configs:      configs,
		S3Repository: s3Repo,
	}
}

//LoadConfig start load config
func (s Service) LoadConfig(ctx context.Context, serviceName string) error {
	tickQuoteCodes, err := s.LoadQuoteCodeData(s.Configs.TickSystem.QuoteCodesDefinitionTickKei1Object)
	if err != nil {
		return err
	}
	isValidQuoteCodes := ValidateUniqueEndpoint(tickQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}
	kehaiQuoteCodes, err := s.LoadQuoteCodeData(s.Configs.TickSystem.QuoteCodesDefinitionKehaiKei1Object)
	if err != nil {
		return err
	}
	isValidQuoteCodes = ValidateUniqueEndpoint(kehaiQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}
	tickQuoteCodes, err = s.LoadQuoteCodeData(s.Configs.TickSystem.QuoteCodesDefinitionTickKei2Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes = ValidateUniqueEndpoint(tickQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	kehaiQuoteCodes, err = s.LoadQuoteCodeData(s.Configs.TickSystem.QuoteCodesDefinitionKehaiKei2Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes = ValidateUniqueEndpoint(kehaiQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	var envKeys, params []string
	var isEnoughParams bool
	runningType := os.Args[2]

	if serviceName == model.BP03 {
		if runningType == model.BP03FirstRunningType {
			envKeys = s.Configs.BP03FirstRunningTypeEnvKeys
		}
		if runningType == model.BP03SecondRunningType {
			envKeys = s.Configs.BP03SecondRunningTypeEnvKeys
		}
	} else if serviceName == model.BP05 {
		switch runningType {
		case model.BP05DemegetRunningType, model.BP05DemegetEvRunningType, model.BP05DemegetPtsRunningType, model.BP05DemegetPtsEvRunningType:
			envKeys = s.Configs.BP05FirstRunningTypeEnvKeys
		case model.BP05DownloadAllRunningType, model.BP05DownloadAllEvoRunningType, model.BP05DownloadAllPtsRunningType, model.BP05DownloadAllEvPtsRunningType:
			envKeys = s.Configs.BP05SecondRunningTypeEnvKeys
		case model.BP05RecreateRunningType, model.BP05RecreatePtsRunningType:
			envKeys = s.Configs.BP05ThirdRunningTypeEnvKeys
		case model.BP05ReDownloadRunningType, model.BP05ReDownloadEvRunningType:
			envKeys = s.Configs.BP05FourthRunningTypeEnvKeys
		}
	} else if serviceName == model.BP06 {
		envKeys = s.Configs.BP06EnvKeyParams
	} else if serviceName == model.BP07 {
		if runningType == model.BP07FirstRunningType {
			envKeys = s.Configs.BP07FirstRunningTypeEnvKeys
		}
		if runningType == model.BP07SecondRunningType {
			envKeys = s.Configs.BP07SecondRunningTypeEnvKeys
		}
		if runningType == model.BP07ThirdRunningType {
			envKeys = s.Configs.BP07ThirdRunningTypeEnvKeys
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

	s.Configs.EnvVarKeys = envKeys
	s.Configs.EnvVarValues = params
	s.Configs.ECS = ecs

	return nil
}

// LoadQuoteCodeData get quote code data and parse to object
func (s *Service) LoadQuoteCodeData(path string) (map[string][]configs.QuoteCodes, error) {
	var quoteCodesData []byte
	var err error

	// Load quote code
	if s.Configs.TickSystem.DevelopEnvironment {
		quoteCodesData, err = os.ReadFile(path)
	} else {
		quoteCodesData, err = s.S3Repository.Download(path)
	}
	if err != nil {
		return nil, err
	}

	var quoteCodes map[string][]configs.QuoteCodes
	err = json.Unmarshal(quoteCodesData, &quoteCodes)
	if err != nil {
		return nil, err
	}

	for _, arrayEndpoints := range quoteCodes {
		for index := range arrayEndpoints {
			if arrayEndpoints[index].Endpoint == model.EmptyString {
				return nil, fmt.Errorf("TKDB_MASTER_ENDPOINT is required, TKLOGIC_ID: %s, QKBN: %s, SNDC: %s",
					arrayEndpoints[index].LogicID, arrayEndpoints[index].QKbn, arrayEndpoints[index].Sndc)
			}
		}
	}

	return quoteCodes, nil
}

// ValidateUniqueEndpoint validate endpoint follow group
func ValidateUniqueEndpoint(quoteCodes map[string][]configs.QuoteCodes) bool {
	validEndpoints := make(map[string]string)
	for _, arrayQuoteCodes := range quoteCodes {
		for index := range arrayQuoteCodes {
			key := arrayQuoteCodes[index].QKbn + model.StrokeCharacter + arrayQuoteCodes[index].Sndc
			value := arrayQuoteCodes[index].Endpoint + model.StrokeCharacter + arrayQuoteCodes[index].DBName
			endpoint, isExists := validEndpoints[key]
			if !isExists {
				validEndpoints[key] = value

				continue
			}

			if value != endpoint {
				return false
			}
		}
	}

	return true
}
