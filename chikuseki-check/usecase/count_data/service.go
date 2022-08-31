package countdata

import (
	"bytes"
	"chikuseki-check/configs"
	"chikuseki-check/infrastructure/repository"
	"chikuseki-check/model"
	"context"
	"fmt"
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

// CountData start count data from database
func (s *service) CountData(ctx context.Context) error {
	buffer := bytes.Buffer{}
	buffer.WriteString(s.config.ZXD)
	buffer.WriteString(model.EnterLineCRLF)

	// process for the first kei
	buffer.WriteString(fmt.Sprintf("KEI1 : %s %s", s.config.Kubun, s.config.Hassin))
	for dataType, prefix := range model.TablePrefix {
		if dataType != model.KehaiDataString {
			dataType = model.TickDataString
		}
		quoteCodeDefinition, isExists := model.QuoteCodesTheFirstKei[s.config.Kubun+model.StrokeCharacter+s.config.Hassin+model.StrokeCharacter+dataType]
		if !isExists {
			continue
		}

		// init database
		err := s.tickDBRepository.InitConnection(ctx, quoteCodeDefinition.Endpoint, quoteCodeDefinition.DBName, model.TheFirstKei, dataType)
		if err != nil {
			return err
		}

		// check table exists\
		isExists = s.tickDBRepository.CheckTableExists(ctx, prefix, s.config.ZXD, s.config.Kubun, s.config.Hassin, quoteCodeDefinition.DBName, model.TheFirstKei, dataType)
		if !isExists {
			continue
		}

		total, err := s.tickDBRepository.CountNumberRecords(ctx, prefix, s.config.ZXD, s.config.Kubun, s.config.Hassin, quoteCodeDefinition.DBName, model.TheFirstKei, dataType)
		if err != nil {
			return err
		}

		buffer.WriteString(model.SpaceString)
		buffer.WriteString(fmt.Sprintf("%s %d", prefix, total))
	}
	buffer.WriteString(model.EnterLineCRLF)

	// process for the second kei
	buffer.WriteString(fmt.Sprintf("KEI2 : %s %s", s.config.Kubun, s.config.Hassin))
	for dataType, prefix := range model.TablePrefix {
		if dataType != model.KehaiDataString {
			dataType = model.TickDataString
		}
		quoteCodeDefinition, isExists := model.QuoteCodesTheSecondKei[s.config.Kubun+model.StrokeCharacter+s.config.Hassin+model.StrokeCharacter+dataType]
		if !isExists {
			continue
		}

		// init database
		err := s.tickDBRepository.InitConnection(ctx, quoteCodeDefinition.Endpoint, quoteCodeDefinition.DBName, model.TheSecondKei, dataType)
		if err != nil {
			return err
		}

		// check table exists\
		isExists = s.tickDBRepository.CheckTableExists(ctx, prefix, s.config.ZXD, s.config.Kubun, s.config.Hassin, quoteCodeDefinition.DBName, model.TheSecondKei, dataType)
		if !isExists {
			continue
		}

		total, err := s.tickDBRepository.CountNumberRecords(ctx, prefix, s.config.ZXD, s.config.Kubun, s.config.Hassin, quoteCodeDefinition.DBName, model.TheSecondKei, dataType)
		if err != nil {
			return err
		}

		buffer.WriteString(model.SpaceString)
		buffer.WriteString(fmt.Sprintf("%s %d", prefix, total))
	}
	// TODO push log to ...
	fmt.Println(buffer.String())

	return nil
}
