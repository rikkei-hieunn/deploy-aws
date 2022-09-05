/*
Package loadconfig implements code two module.
*/
package loadconfig

import (
	"encoding/json"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/infrastructure/repository"
	"os"
)

// Service struct load config environment
type Service struct {
	Config       *configs.Server
	S3Repository repository.IS3Repository
}

// NewService service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) IConfigurationLoader {
	return &Service{
		Config:       cfg,
		S3Repository: s3Repo,
	}
}

// LoadConfig load configuration file
func (s *Service) LoadConfig() error {
	tickDB, calendarDB, err := s.LoadCommonData()
	if err != nil {
		return err
	}
	s.Config.TickDB = *tickDB
	s.Config.TickDB.CalendarEndpoint = *calendarDB

	filebusInfo, err := s.LoadFilebusData()
	if err != nil {
		return err
	}
	s.Config.TickFileBus = *filebusInfo

	return nil
}

// LoadCommonData get common data then parse to object
func (s *Service) LoadCommonData() (*configs.TickDB, *configs.CalendarInfo, error) {
	var commonData []byte
	var err error

	// Load common data
	if s.Config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(s.Config.TickSystem.CommonDefinitionObject)
	} else {
		commonData, err = s.S3Repository.Download(s.Config.TickSystem.CommonDefinitionObject)
	}
	if err != nil {
		return nil, nil, err
	}

	var common struct {
		TickDB     configs.TickDB       `json:"TickDB"`
		CalendarDB configs.CalendarInfo `json:"CalendarInfo"`
	}

	// parse from json content
	err = json.Unmarshal(commonData, &common)
	if err != nil {
		return nil, nil, err
	}

	// validate for TICK Database
	err = common.TickDB.Validate()
	if err != nil {
		return nil, nil, err
	}

	// validate for Calendar Database
	err = common.CalendarDB.Validate()
	if err != nil {
		return nil, nil, err
	}

	return &common.TickDB, &common.CalendarDB, nil
}

// LoadFilebusData get filebus data then parse to object
func (s *Service) LoadFilebusData() (*configs.TickFileBus, error) {
	var filebusData []byte
	var err error

	// Load common data
	if s.Config.TickSystem.DevelopEnvironment {
		filebusData, err = os.ReadFile(s.Config.TickSystem.FilebusDefinitionObject)
	} else {
		filebusData, err = s.S3Repository.Download(s.Config.TickSystem.FilebusDefinitionObject)
	}
	if err != nil {
		return nil, err
	}

	var filebusObj struct {
		Filebus configs.TickFileBus `json:"Filebus"`
	}

	// parse from json content
	err = json.Unmarshal(filebusData, &filebusObj)
	if err != nil {
		return nil, err
	}

	// validate for Filebus object
	err = filebusObj.Filebus.Validate()
	if err != nil {
		return nil, err
	}

	return &filebusObj.Filebus, nil
}
