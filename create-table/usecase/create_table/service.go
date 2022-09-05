package createtable

import (
	"context"
	"create-table/configs"
	"create-table/infrastructure/repository"
	"create-table/model"
	"fmt"
	"time"
)

type service struct {
	config           *configs.TickSystem
	tickDBRepository repository.ITickDBRepository
}

// NewService constructor create a new service
func NewService(cfg *configs.TickSystem, tickDBRepo repository.ITickDBRepository) ITableCreator {
	return &service{
		config:           cfg,
		tickDBRepository: tickDBRepo,
	}
}

// CreateTables start create table process
func (s *service) CreateTables(ctx context.Context, quoteCodes []configs.TargetCreateTable) []error {
	currentDate := time.Now()
	nextWeek := currentDate.AddDate(0, 0, 7)

	var errors []error
	for currentDate.Before(nextWeek) {
		currentDate = currentDate.AddDate(0, 0, 1)
		zxd := currentDate.Format(model.DateFormatWithoutStroke)

		for index := range quoteCodes {
			err := s.tickDBRepository.InitConnection(ctx, quoteCodes[index].Endpoint, quoteCodes[index].DBName)
			if err != nil {
				errors = append(errors, fmt.Errorf("init connection fail, endpoint: %s, database name: %s", quoteCodes[index].Endpoint, quoteCodes[index].DBName))

				continue
			}

			isExists := s.tickDBRepository.CheckTableExists(ctx, quoteCodes[index], zxd)
			if isExists {
				continue
			}

			tableName, err := s.tickDBRepository.CreateTable(ctx, quoteCodes[index], zxd)
			if err != nil {
				errors = append(errors, fmt.Errorf("create table fail, endpoint: %s, database name: %s, table name: %s", quoteCodes[index].Endpoint, quoteCodes[index].DBName, tableName))
			}
		}
	}

	return errors
}
