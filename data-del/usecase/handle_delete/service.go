/*
Package handledel implements logics query receive.
*/
package handledel

import (
	"context"
	"data-del/configs"
	"data-del/infrastructure/repository"
	"data-del/model"
	"github.com/rs/zerolog/log"
	"time"
)

type service struct {
	config           *configs.TickSystem
	tickDBRepository repository.ITickDBRepository
	s3Repository     repository.IS3Repository
}

// NewService constructor
func NewService(
	tickSystem *configs.TickSystem,
	tickDB repository.ITickDBRepository,
	s3Storage repository.IS3Repository,
) IRequestHandler {
	return &service{
		config:           tickSystem,
		tickDBRepository: tickDB,
		s3Repository:     s3Storage,
	}
}

// Start starts service
func (s *service) Start(ctx context.Context, dbType string) error {
	today := time.Now()
	for _, expiredDay := range s.config.ExpiredDays {
		mockDayDelete := today.AddDate(0, 0, -expiredDay.Expire)
		for i := 1; i <= s.config.NumberOfDeletedDays; i++ {
			date := mockDayDelete.AddDate(0, 0, -i).Format(model.FormatDate)
			kubun := expiredDay.QKBN
			kubunInsteadOf, isExisted := model.InsteadOfKubun[kubun]
			if isExisted {
				kubun = kubunInsteadOf
			}
			tables, err := s.tickDBRepository.GetListAvailableTable(ctx, kubun, expiredDay.SNDC, date, dbType)
			if err != nil {
				return err
			}
			if len(tables) == 0 {
				continue
			}
			err = s.deleteHandler(ctx, tables, dbType)
			if err != nil {
				return err
			}
		}
	}

	//delete for case all
	if s.config.ExpiredDaysAll.QKBN == model.EmptyString {
		return nil
	}
	mockDayDelete := today.AddDate(0, 0, -s.config.ExpiredDaysAll.Expire)
	for i := 1; i <= s.config.NumberOfDeletedDays; i++ {
		date := mockDayDelete.AddDate(0, 0, -i).Format(model.FormatDate)
		tables, err := s.tickDBRepository.GetListAvailableTableWithSameDate(ctx, date, dbType)
		if err != nil {
			return err
		}
		if len(tables) == 0 {
			continue
		}
		err = s.deleteHandler(ctx, tables, dbType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) deleteHandler(ctx context.Context, table []string, dbType string) error {
	var err error
	var exists bool
	var s3Path string
	for i := range table {
		if s.config.DevelopEnvironment {
			goto start_delete
		}
		//TODO check s3, change config
		s3Path = table[i]
		exists, err = s.s3Repository.CheckObjectExists(ctx, s3Path)
		if err != nil {
			return err
		}
		if !exists {
			log.Error().Msgf("check s3 object key dont existed, table %s should be backup", table[i])

			continue
		}
	start_delete:
		err = s.tickDBRepository.DropTable(ctx, table[i], dbType)
		if err != nil {
			return err
		}
	}

	return nil
}
