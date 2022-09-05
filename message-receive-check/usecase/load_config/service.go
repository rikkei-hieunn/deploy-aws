/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import (
	"context"
	"encoding/json"
	"message-receive-check/configs"
	"message-receive-check/infrastructure/repository"
	"message-receive-check/model"
	"os"
	"strings"
)

// Service structure define service load configuration data
type Service struct {
	Config       *configs.Server
	S3Repository repository.IS3Repository
}

// NewService Service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) IConfigurationLoader {
	return &Service{
		Config:       cfg,
		S3Repository: s3Repo,
	}
}

// LoadConfig load Config process
func (s *Service) LoadConfig(ctx context.Context) error {
	processNames, groups, err := s.LoadGroupData()
	if err != nil {
		return err
	}
	s.Config.ProcessNames = processNames
	s.Config.Groups = groups

	return nil
}

// LoadGroupData load group data then parse to object
func (s *Service) LoadGroupData() ([]string, []configs.Group, error) {
	var groupData []byte
	var err error

	// Load group data
	if s.Config.DevelopEnvironment {
		groupData, err = os.ReadFile(s.Config.TickSystem.GroupsDefinitionObject)
	} else {
		groupData, err = s.S3Repository.Download(s.Config.TickSystem.GroupsDefinitionObject)
	}
	if err != nil {
		return nil, nil, err
	}

	var groups []configs.Group
	err = json.Unmarshal(groupData, &groups)
	if err != nil {
		return nil, nil, err
	}

	var processMap = make(map[string]bool)
	for i := 0; i < len(groups); i++ {
		dataTypes := strings.Split(groups[i].Types, model.CommaCharacter)
		for j := 0; j < len(dataTypes); j++ {
			if dataTypes[j] != model.MultipleQuoteData {
				processMap[groups[i].LogicGroup] = true
			} else {
				processMap[model.KehaiPrefix+groups[i].LogicGroup] = true
			}
		}
	}

	var processNames []string
	for processName, _ := range processMap {
		processNames = append(processNames, processName)
	}

	return processNames, groups, nil
}
