/*
Package handlerequest implements logics query.
*/
package handlerequest

import (
	"bufio"
	"bytes"
	"context"
	"github.com/rs/zerolog/log"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/infrastructure/repository"
	"insert-calendar-infos/model"
	"strings"
)

// Service default handle_request
type Service struct {
	config            *configs.Server
	tickDBRepository  repository.ITickDBRepository
	filebusRepository repository.IFilebusRepository
	s3Repository      repository.IS3Repository
}

// NewService constructor
func NewService(
	tickSystem *configs.Server,
	tickDBRepo repository.ITickDBRepository,
	fileBusRepo repository.IFilebusRepository,
	s3Storage repository.IS3Repository,
) IRequestHandler {
	return &Service{
		config:            tickSystem,
		tickDBRepository:  tickDBRepo,
		filebusRepository: fileBusRepo,
		s3Repository:      s3Storage,
	}
}

// Start starts service
func (s *Service) Start(ctx context.Context, kei string) error {
	// init transaction
	err := s.tickDBRepository.InitTx(ctx, kei)
	if err != nil {
		log.Error().Msgf("error init transaction with kei %s , error : %s", kei, err.Error())

		return err
	}

	// truncate old data
	truncateQuery := GenQueryTruncateCalendar(s.config.TickDB.CalendarEndpoint.TableName)
	err = s.tickDBRepository.ExecWithTx(ctx, truncateQuery.String(), kei, nil)
	if err != nil {
		// rollback transaction
		s.tickDBRepository.RollbackTx(kei)

		return err
	}

	// download content data
	var path, fileName string
	if kei == model.TheFirstKei {
		path = s.config.TickFileBus.PathCalendar1
		fileName = s.config.TickSystem.CalendarFileName1
	} else {
		path = s.config.TickFileBus.PathCalendar2
		fileName = s.config.TickSystem.CalendarFileName2
	}
	firstContent, err := s.filebusRepository.DownloadFile(path, fileName)
	if err != nil {
		// rollback transaction
		s.tickDBRepository.RollbackTx(kei)

		return err
	}

	// process content file
	err = s.HandleCalendarFile(ctx, firstContent, kei)
	if err != nil {
		// rollback transaction
		s.tickDBRepository.RollbackTx(kei)

		return err
	}

	// commit transaction for the first kei
	s.tickDBRepository.CommitTx(kei)

	return nil
}

// HandleCalendarFile handle content calendar file
func (s *Service) HandleCalendarFile(ctx context.Context, data []byte, kei string) error {
	isColumnsLine := true
	// loop line by line message
	for {
		bufferLength, line, errScan := bufio.ScanLines(data, true)
		if errScan != nil {
			return errScan
		}

		if bufferLength == 0 {
			break
		}

		if isColumnsLine {
			isColumnsLine = false
			if bufferLength <= len(data) {
				data = data[bufferLength:]
			}

			continue
		}

		var args []interface{}
		valuesLine := strings.Split(string(line), model.CommaCharacter)
		for indexValue := range valuesLine {
			var value interface{} = valuesLine[indexValue]
			if value == model.EmptyString {
				value = nil
			}
			args = append(args, value)
		}

		// generate query and push to transaction
		queryString := GenQueryInsertCalendar(s.config.TickDB.CalendarEndpoint.TableName).String()
		err := s.tickDBRepository.ExecWithTx(ctx, queryString, kei, args)
		if err != nil {
			return err
		}

		if bufferLength <= len(data) {
			data = data[bufferLength:]
		}
	}

	return nil
}

// GenQueryInsertCalendar generate query insert calendar
func GenQueryInsertCalendar(tableName string) *bytes.Buffer {
	if tableName == model.EmptyString {
		return nil
	}
	var queryInsert bytes.Buffer

	queryInsert.WriteString(model.FormatInsert)
	queryInsert.WriteString(tableName)
	queryInsert.WriteString(model.FormatFieldInsert)
	queryInsert.WriteString(model.FormatInsertValue)
	queryInsert.WriteString(model.FormatInsertOneField)

	return &queryInsert
}

// GenQueryTruncateCalendar generate query truncate calendar
func GenQueryTruncateCalendar(tableName string) *bytes.Buffer {
	if tableName == model.EmptyString {
		return nil
	}
	var queryTruncate bytes.Buffer

	queryTruncate.WriteString(model.SQLTruncateTable)
	queryTruncate.WriteString(tableName)

	return &queryTruncate
}
