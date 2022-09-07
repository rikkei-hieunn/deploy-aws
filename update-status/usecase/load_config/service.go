/*
Package loadconfig implements logics load config from file.
*/
package loadconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"update-status/configs"
	"update-status/infrastructure/repository"
	"update-status/model"
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
func (s *service) LoadConfig() error {
	// load data database status then push them to model
	databaseStatuses, err := s.LoadDatabaseStatus()
	if err != nil {
		return err
	}
	model.TickDatabaseStatuses = databaseStatuses.Tick
	model.KehaiDatabaseStatuses = databaseStatuses.Kehai

	if s.config.TickSystem.RequestType != model.UpdateStatusTypeDBName && s.config.TickSystem.RequestType != model.UpdateStatusTypeGroupID {
		return nil
	}
	quoteCodes, err := s.LoadQuoteCodeData()
	if err != nil {
		return err
	}
	model.QuoteCodeDefinition = quoteCodes

	return nil
}

// LoadDatabaseStatus get database status from S3 and parse to object
func (s *service) LoadDatabaseStatus() (*configs.GroupDatabaseStatusDefinition, error) {
	var databaseStatusData []byte
	var err error

	// Load database status
	if s.config.DevelopEnvironment {
		databaseStatusData, err = os.ReadFile(s.config.TickSystem.DatabaseStatusDefinitionObject)
	} else {
		databaseStatusData, err = s.s3Repository.Download(s.config.TickSystem.DatabaseStatusDefinitionObject)
	}
	if err != nil {
		return nil, err
	}

	var databaseStatuses configs.GroupDatabaseStatusDefinition
	err = json.Unmarshal(databaseStatusData, &databaseStatuses)
	if err != nil {
		return nil, err
	}

	return &databaseStatuses, nil
}

// LoadQuoteCodeData get quote code data and parse to object
func (s *service) LoadQuoteCodeData() (map[string][]configs.QuoteCodes, error) {
	var path string
	switch s.config.TickSystem.Kei {
	case model.TheFirstKei:
		if s.config.TickSystem.DataType == model.TickData {
			path = s.config.TickSystem.QuoteCodesDefinitionTickKei1Object
		} else if s.config.TickSystem.DataType == model.KehaiData {
			path = s.config.TickSystem.QuoteCodesDefinitionKehaiKei1Object
		}
	case model.TheSecondKei:
		if s.config.TickSystem.DataType == model.TickData {
			path = s.config.TickSystem.QuoteCodesDefinitionTickKei2Object
		} else if s.config.TickSystem.DataType == model.KehaiData {
			path = s.config.TickSystem.QuoteCodesDefinitionKehaiKei2Object
		}
	default:
		return nil, fmt.Errorf("invalid kei")
	}

	var quoteCodesData []byte
	var err error

	// Load quote code
	if s.config.DevelopEnvironment {
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
