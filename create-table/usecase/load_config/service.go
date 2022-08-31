/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import (
	"create-table/configs"
	"create-table/infrastructure/repository"
	"create-table/model"
	"create-table/pkg/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	tickDB, tablePrefix, kubunInstead, oneMinuteTablePrefix, err := s.LoadCommonData()
	if err != nil {
		return err
	}
	s.Config.TickDB = *tickDB
	s.Config.TableNamePrefix = *tablePrefix
	model.KubunsInsteadOf = kubunInstead
	model.OneMinuteTablePrefix = *oneMinuteTablePrefix
	prefix := make(map[string]string)
	for index := range s.Config.TableNamePrefix {
		prefix[strconv.Itoa(s.Config.TableNamePrefix[index].DataType)] = s.Config.TableNamePrefix[index].Prefix
	}
	model.TablePrefix = prefix

	// load group data
	groups, err := s.LoadGroupData()
	if err != nil {
		return err
	}

	// load one minute elements data
	oneMinuteTables, err := s.LoadOneMinuteElementsData()
	if err != nil {
		return err
	}

	// load elements data
	model.Tables, err = s.LoadElementsData(oneMinuteTables)
	if err != nil {
		return err
	}

	// load quote code definition data
	if s.Config.TickSystem.CreateType == model.CreateTypeKei1 || s.Config.TickSystem.CreateType == model.CreateTypeBoth {
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

		targetCreateTables := s.ParseTargetCreateTable(tickQuoteCodes, kehaiQuoteCodes, groups, oneMinuteTables)
		model.TargetCreateTableTheFirstKei = append(model.TargetCreateTableTheFirstKei, targetCreateTables...)
	}

	if s.Config.TickSystem.CreateType == model.CreateTypeKei2 || s.Config.TickSystem.CreateType == model.CreateTypeBoth {
		tickQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionTickKei2Object)
		if err != nil {
			return err
		}

		isValidQuoteCodes := ValidateUniqueEndpoint(tickQuoteCodes)
		if !isValidQuoteCodes {
			return fmt.Errorf("too many endpoint at the same kubun and hassin")
		}

		kehaiQuoteCodes, err := s.LoadQuoteCodeData(s.Config.TickSystem.QuoteCodesDefinitionKehaiKei2Object)
		if err != nil {
			return err
		}

		isValidQuoteCodes = ValidateUniqueEndpoint(kehaiQuoteCodes)
		if !isValidQuoteCodes {
			return fmt.Errorf("too many endpoint at the same kubun and hassin")
		}

		targetCreateTables := s.ParseTargetCreateTable(tickQuoteCodes, kehaiQuoteCodes, groups, oneMinuteTables)
		model.TargetCreateTableTheSecondKei = append(model.TargetCreateTableTheSecondKei, targetCreateTables...)
	}

	return nil
}

// ParseTargetCreateTable convert from quote code map tick and kehai data to target create table array
func (s *Service) ParseTargetCreateTable(tickMap, kehaiMap map[string][]configs.QuoteCodes, groups []configs.Group,
	oneMinuteTables map[string]configs.OneMinuteTableDefinition) []configs.TargetCreateTable {
	var result []configs.TargetCreateTable

	// loop tick map
	for logicGroup, quoteCodes := range tickMap {
		var group *configs.Group
		for groupIndex := range groups {
			if logicGroup == groups[groupIndex].LogicGroup {
				group = &groups[groupIndex]

				break
			}
		}

		if group == nil {
			continue
		}

		for _, dataType := range group.TypesString {
			for index := range quoteCodes {
				_, isExists := oneMinuteTables[quoteCodes[index].QKbn]
				if !utils.Contain(result, quoteCodes[index].QKbn, quoteCodes[index].Sndc, model.OneMinuteData) && isExists {
					result = append(result, configs.TargetCreateTable{
						LogicGroup:  logicGroup,
						QKbn:        quoteCodes[index].QKbn,
						Sndc:        quoteCodes[index].Sndc,
						Endpoint:    quoteCodes[index].Endpoint,
						DBName:      quoteCodes[index].DBName,
						DataType:    model.OneMinuteData,
						TablePrefix: model.OneMinuteTablePrefix,
					})
				}

				if utils.Contain(result, quoteCodes[index].QKbn, quoteCodes[index].Sndc, dataType) {
					continue
				}
				result = append(result, configs.TargetCreateTable{
					LogicGroup:  logicGroup,
					QKbn:        quoteCodes[index].QKbn,
					Sndc:        quoteCodes[index].Sndc,
					Endpoint:    quoteCodes[index].Endpoint,
					DBName:      quoteCodes[index].DBName,
					DataType:    dataType,
					TablePrefix: model.TablePrefix[dataType],
				})
			}
		}
	}

	// loop kehai map
	for logicGroup, quoteCodes := range kehaiMap {
		var group *configs.Group
		for groupIndex := range groups {
			if logicGroup == groups[groupIndex].LogicGroup {
				group = &groups[groupIndex]

				break
			}
		}

		if group == nil {
			continue
		}

		for _, dataType := range group.TypesString {
			for index := range quoteCodes {
				_, isExists := oneMinuteTables[quoteCodes[index].QKbn]
				if !utils.Contain(result, quoteCodes[index].QKbn, quoteCodes[index].Sndc, model.OneMinuteData) && isExists {
					result = append(result, configs.TargetCreateTable{
						LogicGroup:  logicGroup,
						QKbn:        quoteCodes[index].QKbn,
						Sndc:        quoteCodes[index].Sndc,
						Endpoint:    quoteCodes[index].Endpoint,
						DBName:      quoteCodes[index].DBName,
						DataType:    model.OneMinuteData,
						TablePrefix: model.OneMinuteTablePrefix,
					})
				}

				if utils.Contain(result, quoteCodes[index].QKbn, quoteCodes[index].Sndc, dataType) {
					continue
				}
				result = append(result, configs.TargetCreateTable{
					LogicGroup:  logicGroup,
					QKbn:        quoteCodes[index].QKbn,
					Sndc:        quoteCodes[index].Sndc,
					Endpoint:    quoteCodes[index].Endpoint,
					DBName:      quoteCodes[index].DBName,
					DataType:    dataType,
					TablePrefix: model.TablePrefix[dataType],
				})
			}
		}
	}

	return result
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

	return &common.TickDB, &common.TableNamePrefix, common.KubunInsteadOf, &common.OneMinuteCommon.OneMinuteTablePrefix, nil
}

// LoadGroupData get group data then parse to object
func (s *Service) LoadGroupData() ([]configs.Group, error) {
	var groupData []byte
	var err error

	// Load group
	if s.Config.TickSystem.DevelopEnvironment {
		groupData, err = os.ReadFile(s.Config.TickSystem.GroupsDefinitionObject)
	} else {
		groupData, err = s.s3Repository.Download(s.Config.TickSystem.GroupsDefinitionObject)
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

	// convert types string to types array
	for index := range groups {
		types := strings.Split(groups[index].Types, model.CommaCharacter)
		groups[index].TypesString = append(groups[index].TypesString, types...)
	}

	return groups, nil
}

// LoadElementsData get elements data then parse to object
func (s *Service) LoadElementsData(oneMinuteTables map[string]configs.OneMinuteTableDefinition) (map[string]map[string]configs.TableDefinition, error) {
	var elementsData []byte
	var err error

	// Load elements
	if s.Config.TickSystem.DevelopEnvironment {
		elementsData, err = os.ReadFile(s.Config.TickSystem.ElementsDefinitionObject)
	} else {
		elementsData, err = s.s3Repository.Download(s.Config.TickSystem.ElementsDefinitionObject)
	}
	if err != nil {
		return nil, err
	}

	var tables map[string]map[string]configs.TableDefinition
	err = json.Unmarshal(elementsData, &tables)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]configs.TableDefinition)
	for kubun, dataTypesMap := range tables {
		dataTypes := make(map[string]configs.TableDefinition)
		for dataType, table := range dataTypesMap {
			// validate start date
			startDate, err := time.Parse(model.DateFormatWithStroke, table.StartDate)
			if err != nil {
				return nil, err
			}

			if startDate.After(time.Now()) {
				continue
			}

			if table.EndDate == model.EmptyString {
				_, isExists := dataTypes[dataType]
				if !isExists {
					dataTypes[dataType] = table
					result[kubun] = dataTypes
				}

				continue
			}

			endDate, err := time.Parse(model.DateFormatWithStroke, table.EndDate)
			if err != nil {
				return nil, err
			}

			if endDate.After(time.Now()) {
				_, isExists := dataTypes[dataType]
				if !isExists {
					dataTypes[dataType] = table
					result[kubun] = dataTypes
				}
			}
		}
	}

	for kubun, oneMinuteTable := range oneMinuteTables {
		mapDataTypes, isExists := result[kubun]
		if !isExists {
			continue
		}

		mapDataTypes[model.OneMinuteData] = configs.TableDefinition{
			Elements:  oneMinuteTable.Elements,
			StartDate: oneMinuteTable.StartDate,
			EndDate:   oneMinuteTable.EndDate,
		}
		result[kubun] = mapDataTypes
	}

	return result, nil
}

// LoadOneMinuteElementsData get one minute elements data then parse to object
func (s *Service) LoadOneMinuteElementsData() (map[string]configs.OneMinuteTableDefinition, error) {
	var elementsData []byte
	var err error

	// Load elements
	if s.Config.TickSystem.DevelopEnvironment {
		elementsData, err = os.ReadFile(s.Config.TickSystem.OneMinuteDefinitionObject)
	} else {
		elementsData, err = s.s3Repository.Download(s.Config.TickSystem.OneMinuteDefinitionObject)
	}
	if err != nil {
		return nil, err
	}

	var oneMinuteTables []configs.OneMinuteTableDefinition
	err = json.Unmarshal(elementsData, &oneMinuteTables)
	if err != nil {
		return nil, err
	}

	result := make(map[string]configs.OneMinuteTableDefinition)
	for _, table := range oneMinuteTables {
		startDate, err := time.Parse(model.DateFormatWithStroke, table.StartDate)
		if err != nil {
			return nil, err
		}

		if startDate.After(time.Now()) {
			continue
		}

		if table.EndDate == model.EmptyString {
			_, isExists := result[table.TkQkbn]
			if !isExists {
				result[table.TkQkbn] = table
			}

			continue
		}

		endDate, err := time.Parse(model.DateFormatWithStroke, table.EndDate)
		if err != nil {
			return nil, err
		}

		if endDate.After(time.Now()) {
			_, isExists := result[table.TkQkbn]
			if !isExists {
				result[table.TkQkbn] = table
			}
		}
	}

	return result, nil
}

// LoadQuoteCodeData get quote code data and parse to object
func (s *Service) LoadQuoteCodeData(path string) (map[string][]configs.QuoteCodes, error) {
	var quoteCodesData []byte
	var err error

	// Load quote code
	if s.Config.TickSystem.DevelopEnvironment {
		quoteCodesData, err = os.ReadFile(path)
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
