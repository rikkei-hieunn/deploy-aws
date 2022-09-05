package updatestatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"update-status/configs"
	"update-status/infrastructure/repository"
	"update-status/model"
)

type serviceUpdateTypeQuoteCode struct {
	config       *configs.TickSystem
	s3Repository repository.IS3Repository
}

// NewUpdateTypeQuoteCodeService constructor init new service update status database
func NewUpdateTypeQuoteCodeService(cfg *configs.TickSystem, s3Repo repository.IS3Repository) IUpdater {
	return &serviceUpdateTypeQuoteCode{
		config:       cfg,
		s3Repository: s3Repo,
	}
}

// StartUpdateStatus start update status database
func (s *serviceUpdateTypeQuoteCode) StartUpdateStatus() error {
	// set new status
	groupNewStatus, err := s.SetNewStatus()
	if err != nil {
		return err
	}

	// parse object to json
	data, err := json.MarshalIndent(groupNewStatus, model.EmptyString, model.TabString)
	if err != nil {
		return err
	}

	// push new status to S3
	if s.config.DevelopEnvironment {
		return ioutil.WriteFile(s.config.DatabaseStatusDefinitionObject, data, 0644)
	}

	return s.s3Repository.Upload(s.config.DatabaseStatusDefinitionObject, data)
}

// SetNewStatus get list old status then update to new status
func (s *serviceUpdateTypeQuoteCode) SetNewStatus() (*configs.GroupDatabaseStatusDefinition, error) {
	var groupStatus configs.GroupDatabaseStatusDefinition

	// get old status from model
	var oldStatuses configs.ArrayDatabaseStatus
	if s.config.DataType == model.TickData {
		oldStatuses = model.TickDatabaseStatuses
		groupStatus.Tick = oldStatuses
		groupStatus.Kehai = model.KehaiDatabaseStatuses
	} else {
		oldStatuses = model.KehaiDatabaseStatuses
		groupStatus.Tick = model.TickDatabaseStatuses
		groupStatus.Kehai = oldStatuses
	}

	// parse request object
	request, ok := s.config.Request.(model.UpdateTypeQuoteCode)
	if !ok {
		return nil, fmt.Errorf("invalid request update status")
	}

	isValid := false
	// loop old statuses and check kubun, hassin then update new status
	for index := range oldStatuses {
		if oldStatuses[index].QKbn != request.Kubun || oldStatuses[index].Sndc != request.Hassin {
			continue
		}

		isValid = true
		if s.config.Kei == model.TheFirstKei {
			oldStatuses[index].TheFirstKeiStatus = request.NewStatus
		} else {
			oldStatuses[index].TheSecondKeiStatus = request.NewStatus
		}
	}

	if !isValid {
		return nil, fmt.Errorf("invalid kubun and hassin")
	}

	return &groupStatus, nil
}
