package loadconfig

import (
	"bytes"
	"context"
	"data-del/configs"
	"data-del/infrastructure/repository"
	"data-del/model"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

//LoadConfig start load configuration
func (s *Service) LoadConfig(ctx context.Context, keiNumber string) error {
	var err error
	var dataObject *bytes.Buffer
	var tickEndpointPath, kehaiEndpointPath string
	var commonData, tickEndpointData, kehaiEndpointData, expireData []byte

	// Load common data
	if s.Config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(s.Config.TickSystem.CommonDefinitionObject)
		if err != nil {
			return err
		}
	} else {
		dataObject, err = s.S3Repository.GetObject(ctx, s.Config.TickSystem.CommonDefinitionObject)
		if err != nil {
			return err
		}
		commonData = dataObject.Bytes()
	}
	tablePrefix, tickDB, InsteadOfKubun, err := s.ParseCommonData(commonData)
	if err != nil {
		return err
	}
	s.Config.TablePrefix = tablePrefix
	s.Config.TickDB = *tickDB
	model.InsteadOfKubun = InsteadOfKubun
	err = s.Config.TickDB.Validate()
	if err != nil {
		return err
	}

	// Load expire data
	if s.Config.TickSystem.DevelopEnvironment {
		expireData, err = os.ReadFile(filepath.Clean(s.Config.TickSystem.ExpireDefinitionObject))
		if err != nil {
			return err
		}
	} else {
		dataObject, err = s.S3Repository.GetObject(ctx, s.Config.TickSystem.ExpireDefinitionObject)
		if err != nil {
			return err
		}
		expireData = dataObject.Bytes()
	}

	expiredDays, expiredDayAll, err := s.ParseExpireData(expireData)
	if err != nil {
		return fmt.Errorf("parse expired data fail %w", err)
	}
	s.Config.ExpiredDays = expiredDays
	s.Config.ExpiredDaysAll = *expiredDayAll

	if keiNumber == model.FirstKei {
		tickEndpointPath = s.Config.TickDB1EndpointDefinitionObject
		kehaiEndpointPath = s.Config.KehaiDB1EndpointDefinitionObject
	} else {
		tickEndpointPath = s.Config.TickDB2EndpointDefinitionObject
		kehaiEndpointPath = s.Config.KehaiDB2EndpointDefinitionObject
	}
	endpoints := make(map[string]map[string][]string)

	if s.Config.TickSystem.DevelopEnvironment {
		tickEndpointData, err = os.ReadFile(filepath.Clean(tickEndpointPath))
		if err != nil {
			return err
		}
		kehaiEndpointData, err = os.ReadFile(filepath.Clean(kehaiEndpointPath))
		if err != nil {
			return err
		}
	} else {
		buff, err := s.S3Repository.GetObject(ctx, tickEndpointPath)
		if err != nil {
			return err
		}
		tickEndpointData = buff.Bytes()

		buff, err = s.S3Repository.GetObject(ctx, kehaiEndpointPath)
		if err != nil {
			return err
		}
		kehaiEndpointData = buff.Bytes()
	}
	endpointTick, err := s.ParseDBEndpointData(tickEndpointData)
	if err != nil {
		return err
	}
	endpointKehai, err := s.ParseDBEndpointData(kehaiEndpointData)
	if err != nil {
		return err
	}
	endpoints[model.TypeKehai] = endpointKehai
	endpoints[model.TypeTick] = endpointTick

	s.Config.Endpoints = endpoints

	return nil
}

//ParseCommonData parse common data : db configuration and table prefix
func (s *Service) ParseCommonData(data []byte) (map[int]string, *configs.TickDB, map[string]string, error) {
	var common struct {
		TickDB          configs.TickDB            `json:"TickDB"`
		TableNamePrefix []configs.TickTablePrefix `json:"TableNamePrefix"`
		InsteadOfKubun  map[string]string         `json:"InsteadOfKubun"`
	}
	if err := json.Unmarshal(data, &common); err != nil {
		return nil, nil, nil, err
	}
	result := make(map[int]string)
	for _, tablePrefix := range common.TableNamePrefix {
		result[tablePrefix.DataType] = tablePrefix.Prefix
	}

	return result, &common.TickDB, common.InsteadOfKubun, nil
}

//ParseExpireData parse expired days return list of expired data and one item expired day all
func (s *Service) ParseExpireData(data []byte) (expiredDay []configs.TickExpire, expiredAll *configs.TickExpire, err error) {
	var res []configs.TickExpire
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, nil, err
	}
	for i := range res {
		if res[i].QKBN == model.QKBNAll {
			expiredAll = &res[i]

			continue
		}
		expiredDay = append(expiredDay, res[i])
	}

	return expiredDay, expiredAll, nil
}

//ParseDBEndpointData parse endpoint with key is endpoint/dbName values is a list of kubun/hassin
func (s *Service) ParseDBEndpointData(data []byte) (map[string][]string, error) {
	inputEndPointMap := make(map[string][]configs.EndPoint, 0)

	if err := json.Unmarshal(data, &inputEndPointMap); err != nil {
		return nil, err
	}
	results := make(map[string][]string)
	for _, endpoints := range inputEndPointMap {
		for _, endpoint := range endpoints {
			key := endpoint.DBEndpoint + model.StrokeCharacter + endpoint.DBName
			kubunHasshinPair := endpoint.TKQKBN + model.StrokeCharacter + endpoint.SNDC
			results[key] = append(results[key], kubunHasshinPair)
		}
	}

	return results, nil
}
