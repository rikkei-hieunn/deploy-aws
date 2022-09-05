/*
Package loadconfig implements logics load config from file.
*/
package loadconfig

import (
	"context"
	"ec2-all-stop/configs"
	"ec2-all-stop/infrastructure/repository"
	"ec2-all-stop/model"
	"encoding/json"
	"fmt"
	"os"
)

type service struct {
	config       *configs.Server
	s3Repository repository.IS3Repository
}

// NewService service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) IConfigurationLoader {
	return &service{
		config:       cfg,
		s3Repository: s3Repo,
	}
}

// LoadConfig load config process
func (s *service) LoadConfig(ctx context.Context) error {
	var err error
	var commonData []byte
	var groupDefinitionObject string
	instanceIDMap := make(map[string]string)

	if s.config.TickSystem.Kei == model.TheFirstKei {
		groupDefinitionObject = s.config.TickSystem.GroupsDefinitionKei1Object
	} else {
		groupDefinitionObject = s.config.TickSystem.GroupsDefinitionKei2Object
	}

	// Load group definition data
	if s.config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(groupDefinitionObject)
	} else {
		commonData, err = s.s3Repository.Download(groupDefinitionObject)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(commonData, &s.config.HostNameDefinitions)
	if err != nil {
		return err
	}
	for _, hostNameDefinition := range s.config.HostNameDefinitions {
		if hostNameDefinition.TickHostName != model.EmptyString {
			instanceIDMap[hostNameDefinition.TickHostName] = os.Getenv(hostNameDefinition.TickHostName + model.MachineSuffix)
		}
		if hostNameDefinition.KehaiHostName != model.EmptyString {
			instanceIDMap[hostNameDefinition.KehaiHostName] = os.Getenv(hostNameDefinition.KehaiHostName + model.MachineSuffix)
		}
	}
	for _, id := range instanceIDMap {
		if id == model.EmptyString {
			continue
		}
		s.config.InstanceIds = append(s.config.InstanceIds, id)
	}

	if len(s.config.InstanceIds) == 0 {
		return fmt.Errorf("list instance's id are empty")
	}

	return err
}
