/*
Package updatestatus implement logic update status database
*/
package updatestatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"update-status/configs"
	"update-status/infrastructure/repository"
	"update-status/model"
)

type serviceUpdateTypeGroupID struct {
	config       *configs.TickSystem
	s3Repository repository.IS3Repository
}

// NewUpdateTypeGroupIDService constructor init new service update status database
func NewUpdateTypeGroupIDService(cfg *configs.TickSystem, s3Repo repository.IS3Repository) IUpdater {
	return &serviceUpdateTypeGroupID{
		config:       cfg,
		s3Repository: s3Repo,
	}
}

// StartUpdateStatus start update status database
func (s *serviceUpdateTypeGroupID) StartUpdateStatus() error {
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
func (s *serviceUpdateTypeGroupID) SetNewStatus() (*configs.GroupDatabaseStatusDefinition, error) {
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
	request, ok := s.config.Request.(model.UpdateTypeGroupID)
	if !ok {
		return nil, fmt.Errorf("invalid request update status")
	}

	groupData, isExists := model.QuoteCodeDefinition[request.GroupID]
	if !isExists {
		return nil, fmt.Errorf("group id not found")
	}

	// loop old statuses and check kubun, hassin then update new status
	for index := range groupData {
		kubun := groupData[index].QKbn
		hassin := groupData[index].Sndc

		for statusIndex := range oldStatuses {
			if oldStatuses[statusIndex].QKbn != kubun || oldStatuses[statusIndex].Sndc != hassin {
				continue
			}

			if s.config.Kei == model.TheFirstKei {
				oldStatuses[statusIndex].TheFirstKeiStatus = request.NewStatus
			} else {
				oldStatuses[statusIndex].TheSecondKeiStatus = request.NewStatus
			}
		}
	}

	return &groupStatus, nil
}
