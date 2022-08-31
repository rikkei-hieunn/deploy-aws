/*
Package loadconfig implements code two module.
*/
package loadconfig

import (
	"context"
	"encoding/json"
	"os"
	"recreate-one-minute/configs"
	"recreate-one-minute/infrastructure/repository"
	"recreate-one-minute/model"
	"recreate-one-minute/pkg"
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

// LoadConfig load configuration environment_variables.json file
func (s *service) LoadConfig(ctx context.Context) error {
	tickDB, candleManagementPrefix, kubunInsteadOf, err := s.LoadCommonData()
	if err != nil {
		return err
	}
	s.config.TickDB = *tickDB
	s.config.TickSystem.CandleTablePrefix = *candleManagementPrefix
	model.KubunInsteadOf = kubunInsteadOf

	endpointMap, err := s.LoadEndpointData()
	if err != nil {
		return err
	}
	s.config.TickDB.Endpoints = endpointMap

	return nil
}

// LoadCommonData load common data then parse to object
func (s *service) LoadCommonData() (*configs.TickDB, *string, map[string]string, error) {
	var commonData []byte
	var err error

	// Load common data
	if s.config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(s.config.TickSystem.CommonDefinitionObject)
	} else {
		commonData, err = s.s3Repository.Download(s.config.TickSystem.CommonDefinitionObject)
	}
	if err != nil {
		return nil, nil, nil, err
	}

	var common struct {
		TickDB          configs.TickDB          `json:"TickDB"`
		KubunInsteadOf  map[string]string       `json:"InsteadOfKubun"`
		OneMinuteCommon configs.OneMinuteCommon `json:"OneMinuteCommon"`
	}

	// parse from json content
	err = json.Unmarshal(commonData, &common)
	if err != nil {
		return nil, nil, nil, err
	}

	err = common.TickDB.Validate()
	if err != nil {
		return nil, nil, nil, err
	}

	return &common.TickDB, &common.OneMinuteCommon.CandleManagementTablePrefix, common.KubunInsteadOf, nil
}

// LoadEndpointData load endpoint data then parse to object
func (s *service) LoadEndpointData() (map[string][]string, error) {
	var endpointObject string
	if s.config.TickSystem.Kei == model.TheFirstKei {
		endpointObject = s.config.TickSystem.DB1EndpointDefinitionObject
	} else {
		endpointObject = s.config.TickSystem.DB2EndpointDefinitionObject
	}

	// Load common data
	var err error
	var endpointData []byte
	if s.config.TickSystem.DevelopEnvironment {
		endpointData, err = os.ReadFile(endpointObject)
	} else {
		endpointData, err = s.s3Repository.Download(endpointObject)
	}
	if err != nil {
		return nil, err
	}

	var endpointMap map[string][]configs.EndPoint
	err = json.Unmarshal(endpointData, &endpointMap)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, endpoint := range endpointMap {
		for _, item := range endpoint {
			kubunHasshin := item.TKQKBN + model.StrokeCharacter + item.SNDC
			key := item.DBEndpoint + model.StrokeCharacter + item.DBName
			if pkg.IsItemExisted(kubunHasshin, result[key]) {
				continue
			}
			result[key] = append(result[key], kubunHasshin)
		}
	}

	return result, nil
}
