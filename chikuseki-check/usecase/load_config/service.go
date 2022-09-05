/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import (
	"chikuseki-check/configs"
	"chikuseki-check/infrastructure/repository"
	"chikuseki-check/model"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Service structure service load Config data
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
func (s *Service) LoadConfig() error {
	// load common data then put them to Config object and model
	tickDB, tablePrefix, kubunInstead, err := s.LoadCommonData()
	if err != nil {
		return err
	}
	s.Config.TickDB = *tickDB
	s.Config.TableNamePrefix = *tablePrefix
	model.KubunsInsteadOf = kubunInstead
	prefix := make(map[string]string)
	for index := range s.Config.TableNamePrefix {
		if s.Config.TableNamePrefix[index].DataType == model.HigaData {
			continue
		}
		prefix[strconv.Itoa(s.Config.TableNamePrefix[index].DataType)] = s.Config.TableNamePrefix[index].Prefix
	}
	model.TablePrefix = prefix

	// load quote code definition data
	// load for the first kei
	tickQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionTickKei1Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes := ValidateUniqueEndpoint(tickQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	kehaiQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionKehaiKei1Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes = ValidateUniqueEndpoint(kehaiQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	validQuoteCodeFirstKei := s.GetQuoteCodes(s.Config.TickSystem.Kubun, s.Config.TickSystem.Hassin, tickQuoteCodes, kehaiQuoteCodes)
	if len(validQuoteCodeFirstKei) == 0 {
		return fmt.Errorf("kubun and hassin not found in the first kei")
	}

	// load for the second kei
	tickQuoteCodes, err = s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionTickKei2Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes = ValidateUniqueEndpoint(tickQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	kehaiQuoteCodes, err = s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionKehaiKei2Object)
	if err != nil {
		return err
	}

	isValidQuoteCodes = ValidateUniqueEndpoint(kehaiQuoteCodes)
	if !isValidQuoteCodes {
		return fmt.Errorf("too many endpoint at the same kubun and hassin")
	}

	validQuoteCodeSecondKei := s.GetQuoteCodes(s.Config.TickSystem.Kubun, s.Config.TickSystem.Hassin, tickQuoteCodes, kehaiQuoteCodes)
	if len(validQuoteCodeSecondKei) == 0 {
		return fmt.Errorf("kubun and hassin not found in the second kei")
	}
	model.QuoteCodesTheFirstKei = validQuoteCodeFirstKei
	model.QuoteCodesTheSecondKei = validQuoteCodeSecondKei

	return nil
}

// LoadCommonData get common data then parse to object
func (s *Service) LoadCommonData() (*configs.TickDB, *configs.TableNamePrefix, map[string]string, error) {
	var commonData []byte
	var err error

	// Load common data
	if s.Config.TickSystem.DevelopEnvironment {
		commonData, err = os.ReadFile(s.Config.TickSystem.CommonDefinitionObject)
	} else {
		commonData, err = s.S3Repository.Download(s.Config.TickSystem.CommonDefinitionObject)
	}
	if err != nil {
		return nil, nil, nil, err
	}

	var common struct {
		TickDB          configs.TickDB          `json:"TickDB"`
		TableNamePrefix configs.TableNamePrefix `json:"TableNamePrefix"`
		KubunInsteadOf  map[string]string       `json:"InsteadOfKubun"`
	}

	// parse from json content
	err = json.Unmarshal(commonData, &common)
	if err != nil {
		return nil, nil, nil, err
	}

	// validate for TICK Database
	err = common.TickDB.Validate()
	if err != nil {
		return nil, nil, nil, err
	}

	return &common.TickDB, &common.TableNamePrefix, common.KubunInsteadOf, nil
}

// GetQuoteCodes get quote code definition follow kubun and hassin
func (s *Service) GetQuoteCodes(kubun, hassin string, tickQuoteCodes, kehaiQuoteCodes map[string][]configs.QuoteCodes) map[string]configs.QuoteCodes {
	quoteCodesMap := make(map[string]configs.QuoteCodes)
	for _, arrayQuoteCodes := range tickQuoteCodes {
		for index := range arrayQuoteCodes {
			if arrayQuoteCodes[index].QKbn != kubun || arrayQuoteCodes[index].Sndc != hassin {
				continue
			}

			key := arrayQuoteCodes[index].QKbn + model.StrokeCharacter + arrayQuoteCodes[index].Sndc + model.StrokeCharacter + strconv.Itoa(model.TickData)
			quoteCodesMap[key] = arrayQuoteCodes[index]

			break
		}
	}

	for _, arrayQuoteCodes := range kehaiQuoteCodes {
		for index := range arrayQuoteCodes {
			if arrayQuoteCodes[index].QKbn != kubun || arrayQuoteCodes[index].Sndc != hassin {
				continue
			}

			key := arrayQuoteCodes[index].QKbn + model.StrokeCharacter + arrayQuoteCodes[index].Sndc + model.StrokeCharacter + strconv.Itoa(model.KehaiData)
			quoteCodesMap[key] = arrayQuoteCodes[index]

			break
		}
	}

	return quoteCodesMap
}

// LoadQuoteCodeData get quote code data and parse to object
func (s *Service) LoadQuoteCodeData(path string) (map[string][]configs.QuoteCodes, error) {
	var quoteCodesData []byte
	var err error

	// Load quote code
	if s.Config.TickSystem.DevelopEnvironment {
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
