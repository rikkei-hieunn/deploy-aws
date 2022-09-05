/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
	"process-get-data/configs"
	"process-get-data/infrastructure/repository"
	"process-get-data/model"
	"strconv"
)

type Service struct {
	Config       *configs.Server
	s3Repository repository.IS3Repository
}

// NewService Service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) IConfigurationLoader {
	return &Service{
		Config:       cfg,
		s3Repository: s3Repo,
	}
}

// LoadConfig load Config process
func (s *Service) LoadConfig() error {
	// load common data then put them to Config object and model
	tickDB, tablePrefix, kubunInstead, candleManagementPrefix, err := s.LoadCommonData()
	if err != nil {
		return err
	}
	s.Config.TickDB = *tickDB
	s.Config.TableNamePrefix = *tablePrefix
	model.KubunsInsteadOf = kubunInstead
	model.CandleManagementPrefix = *candleManagementPrefix
	prefix := make(map[string]string)
	for index := range s.Config.TableNamePrefix {
		prefix[strconv.Itoa(s.Config.TableNamePrefix[index].DataType)] = s.Config.TableNamePrefix[index].Prefix
	}
	model.TablePrefix = prefix

	// load one minute Config
	oneMinuteConfigs, err := s.LoadOneMinuteConfigData()
	if err != nil {
		return err
	}
	model.OneMinuteConfigs = oneMinuteConfigs

	// load quote code definition data
	if s.Config.TickSystem.ProcessGetData == model.GetDataTypeKei1 || s.Config.TickSystem.ProcessGetData == model.GetDataTypeBoth {
		tickQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionTickKei1Object)
		if err != nil {
			return err
		}
		model.QuoteCodesDefinitionTheFirstKei = tickQuoteCodes
	}

	if s.Config.TickSystem.ProcessGetData == model.GetDataTypeKei2 || s.Config.TickSystem.ProcessGetData == model.GetDataTypeBoth {
		tickQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionTickKei2Object)
		if err != nil {
			return err
		}
		model.QuoteCodesDefinitionTheSecondKei = tickQuoteCodes
	}

	return nil
}

// LoadCommonData get common data then parse to object
func (s *Service) LoadCommonData() (*configs.TickDB, *configs.TableNamePrefix, map[string]string, *string, error) {
	var commonData []byte
	var err error

	// Load common data
	if s.Config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(s.Config.TickSystem.CommonDefinitionObject)
	} else {
		commonData, err = s.s3Repository.Download(s.Config.TickSystem.CommonDefinitionObject)
	}
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var common struct {
		TickDB          configs.TickDB          `json:"TickDB"`
		TableNamePrefix configs.TableNamePrefix `json:"TableNamePrefix"`
		KubunInsteadOf  map[string]string       `json:"InsteadOfKubun"`
		OneMinuteCommon configs.OneMinuteCommon `json:"OneMinuteCommon"`
	}

	// parse from json content
	err = json.Unmarshal(commonData, &common)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// validate for TICK Database
	err = common.TickDB.Validate()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &common.TickDB, &common.TableNamePrefix, common.KubunInsteadOf, &common.OneMinuteCommon.CandleManagementTablePrefix, nil
}

// LoadQuoteCodeData get quote code data and parse to object
func (s *Service) LoadQuoteCodeData(path string) (map[string]configs.QuoteCodes, error) {
	var quoteCodesData []byte
	var err error

	// Load quote code
	if s.Config.TickSystem.DevelopEnvironment {
		quoteCodesData, err = os.ReadFile(filepath.Clean(path))
	} else {
		quoteCodesData, err = s.s3Repository.Download(path)
	}
	if err != nil {
		return nil, err
	}

	var quoteCodes map[string][]configs.QuoteCodes
	err = json.Unmarshal(quoteCodesData, &quoteCodes)
	if err != nil {
		return nil, err
	}

	quoteCodesMap := make(map[string]configs.QuoteCodes)
	for _, arrayQuoteCodes := range quoteCodes {
		for index := range arrayQuoteCodes {
			key := arrayQuoteCodes[index].QKbn + model.StrokeCharacter + arrayQuoteCodes[index].Sndc
			quoteCodesMap[key] = arrayQuoteCodes[index]
		}
	}

	return quoteCodesMap, nil
}

// LoadOneMinuteConfigData get one minute Config data and parse to object
func (s *Service) LoadOneMinuteConfigData() ([]configs.OneMinuteConfig, error) {
	var oneMinuteConfigData []byte
	var err error

	// Load quote code
	if s.Config.TickSystem.DevelopEnvironment {
		oneMinuteConfigData, err = os.ReadFile(s.Config.TickSystem.OneMinuteOperatorConfigObject)
	} else {
		oneMinuteConfigData, err = s.s3Repository.Download(s.Config.TickSystem.OneMinuteOperatorConfigObject)
	}
	if err != nil {
		return nil, err
	}

	var oneMinuteConfigs []configs.OneMinuteConfig
	err = json.Unmarshal(oneMinuteConfigData, &oneMinuteConfigs)
	if err != nil {
		return nil, err
	}

	for index := range oneMinuteConfigs {
		oneMinuteConfigs[index].OriginStartIndex = oneMinuteConfigs[index].StartIndex
		if oneMinuteConfigs[index].StartIndex == model.EmptyString {
			oneMinuteConfigs[index].StartIndex = model.DefaultStartIndex
		}
	}

	return oneMinuteConfigs, nil
}
