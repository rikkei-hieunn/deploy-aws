package sendrequest

import (
	"encoding/json"
	"fmt"
	"os"
	"send-command/configs"
	"send-command/infrastructure/repository"
	"send-command/model"
	"strconv"
)

type serviceSendRequestGroup struct {
	config           *configs.TickSystem
	s3Repository     repository.IS3Repository
	socketRepository repository.ISocketRepository
}

// NewSendGroupService constructor init service send command
func NewSendGroupService(cfg *configs.TickSystem, s3Repo repository.IS3Repository, socketRepo repository.ISocketRepository) ISender {
	return &serviceSendRequestGroup{
		config:           cfg,
		s3Repository:     s3Repo,
		socketRepository: socketRepo,
	}
}

// HandleRequest handle request and send command
func (s *serviceSendRequestGroup) HandleRequest() error {
	request, ok := s.config.Request.(model.RequestGroupID)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	// check kei and download group definition from S3
	var groupDefinition string
	if request.Kei == model.TheFirstKei {
		groupDefinition = s.config.GroupsDefinitionKei1Object
	} else {
		groupDefinition = s.config.GroupsDefinitionKei2Object
	}

	var err error
	var data []byte
	if s.config.DevelopEnvironment {
		data, err = os.ReadFile(groupDefinition)
	} else {
		data, err = s.s3Repository.Download(groupDefinition)
	}
	if err != nil {
		return err
	}

	// parse list groups definition
	var groups []model.GroupDefinition
	err = json.Unmarshal(data, &groups)
	if err != nil {
		return err
	}

	// loop groups array end send command
	for _, group := range groups {
		if group.GroupID != request.GroupID {
			continue
		}

		portNumber, err := strconv.Atoi(group.Port)
		if err != nil {
			return fmt.Errorf("invalid port number, group ID: %s", group.GroupID)
		}

		if request.DataType == model.TickData && group.TickMachineName != model.EmptyString {
			err = s.socketRepository.InitWriter(group.TickMachineName, portNumber)
			if err != nil {
				return err
			}

			err = s.socketRepository.SendCommand(request.Command)
			if err != nil {
				return err
			}
		}
		if request.DataType == model.KehaiData && group.KehaiMachineName != model.EmptyString {
			err = s.socketRepository.InitWriter(group.KehaiMachineName, portNumber)
			if err != nil {
				return err
			}

			err = s.socketRepository.SendCommand(request.Command)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
