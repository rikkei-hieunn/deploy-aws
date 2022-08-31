package loadconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"start-jushin/configs"
	"start-jushin/infrastructure/repository"
	"start-jushin/model"
	"strings"
)

type Service struct {
	Configs      *configs.Server
	S3Repository repository.IS3Repository
}

//NewService load config constructor
func NewService(cfgs *configs.Server, s3Repo repository.IS3Repository) ILoadConfig {
	return &Service{
		Configs:      cfgs,
		S3Repository: s3Repo,
	}
}

//LoadConfig load config function
func (s *Service) LoadConfig(startType, keiType, groupID, dataType, groupLine string) error {
	var err error
	var path string
	var groups []configs.Group
	var instanceIds []string

	if keiType == model.FirstKei {
		path = s.Configs.GroupDefinitionForFirstKei
	} else if keiType == model.SecondKei {
		path = s.Configs.GroupDefinitionForSecondKei
	}

	if dataType != model.EmptyString && dataType != model.TypeKehai && dataType != model.TypeTick {
		return fmt.Errorf(" data type : %s is invalid", dataType)
	}

	groups, err = s.LoadGroupData(path)
	if err != nil {
		return err
	}

	ssm := make(map[string]string)
	if startType == model.TypeRunAll {

		for i := range groups {
			err = groups[i].Validate()
			if err != nil {
				return err
			}
			err = s.parseInstanceIdsAllHost(&groups[i], ssm)
			if err != nil {
				return err
			}
		}
	} else if startType == model.TypeRunSSS {
		for i := range groups {
			if groups[i].GroupID != groupID {
				continue
			}

			err = groups[i].Validate()
			if err != nil {
				return err
			}

			err = s.parseInstanceIdsSpecificHost(dataType, &groups[i], ssm)
			if err != nil {
				return err
			}
		}
	} else if startType == model.TypeRunByGroupLine {
		for i := range groups {
			if groups[i].GroupLine != groupLine {
				continue
			}
			err = groups[i].Validate()
			if err != nil {
				return err
			}
			err = s.parseInstanceIdsAllHost(&groups[i], ssm)
			if err != nil {
				return err
			}
		}
	}

	for _, id := range ssm {
		instanceIds = append(instanceIds, id)
	}

	s.Configs.InstancesIDs = instanceIds
	folderPath := os.Getenv(s.Configs.InstancePathKey)
	if folderPath == model.EmptyString {
		return fmt.Errorf("folder path is not found")
	}
	lastIndex := strings.LastIndex(folderPath, model.Stroke)
	folder := folderPath[:lastIndex]
	program := folderPath[lastIndex:]

	s.Configs.Commands = append(s.Configs.Commands, "cd "+folder, "."+program)

	return nil
}

// LoadGroupData get group data then parse to object
func (s *Service) LoadGroupData(path string) ([]configs.Group, error) {
	var groupData []byte
	var err error

	// Load group
	if s.Configs.TickSystem.DevelopEnvironment {
		groupData, err = os.ReadFile(path)
	} else {
		groupData, err = s.S3Repository.Download(path)
	}
	if err != nil {
		return nil, err
	}

	var groups []configs.Group
	// parse from json content
	err = json.Unmarshal(groupData, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

//parseInstanceIdsAllHost parse instance id with two host tick and kehai
func (s *Service) parseInstanceIdsAllHost(group *configs.Group, ssm map[string]string) error {
	if _, isExisted := ssm[group.KehaiHostName]; !isExisted {
		instanceID := os.Getenv(group.KehaiHostName + model.MachineSuffix)
		if instanceID != model.EmptyString {
			ssm[group.KehaiHostName] = instanceID
		}
	}
	if _, isExisted := ssm[group.TickHostName]; !isExisted {
		instanceID := os.Getenv(group.TickHostName + model.MachineSuffix)
		if instanceID != model.EmptyString {
			ssm[group.TickHostName] = instanceID
		}
	}

	return nil
}

func (s *Service) parseInstanceIdsSpecificHost(dataType string, group *configs.Group, ssm map[string]string) error {
	var hostName string

	if dataType == model.TypeTick {
		hostName = group.TickHostName
	} else {
		hostName = group.KehaiHostName
	}
	if _, isExisted := ssm[hostName]; !isExisted {
		instanceID := os.Getenv(hostName + model.MachineSuffix)
		if instanceID != model.EmptyString {
			ssm[hostName] = instanceID
		}
	}

	return nil
}
