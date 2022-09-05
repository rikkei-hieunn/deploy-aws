/*
Package handlerecreate implements logics about receive message.
*/
package handlerecreate

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
	"recreate-one-minute/configs"
	"recreate-one-minute/infrastructure/repository"
	"recreate-one-minute/model"
	"recreate-one-minute/pkg"
)

// service default handle_request
type service struct {
	DBRepository repository.ITickDBRepository
	Cfg          *configs.Server
}

// NewService constructor
func NewService(
	cfg *configs.Server,
	dbRepository repository.ITickDBRepository,
) IRequestHandler {
	return &service{
		dbRepository,
		cfg,
	}
}

// Start handle logic application
func (s *service) Start(ctx context.Context) error {
	var validRecord []model.Record
	var err error
	for _, kubunHasshinPairs := range s.Cfg.Endpoints {
		if len(kubunHasshinPairs) == 0 {
			continue
		}
		for _, pair := range kubunHasshinPairs {
			table := pkg.GetTableName(s.Cfg.CandleTablePrefix, pair)
			if table == model.EmptyString {
				continue
			}
			args := []interface{}{
				table,
				model.StatusFail,
			}
			err = s.DBRepository.GetDataFromCandleManagement(ctx, pair, args, &validRecord)
			if err != nil {
				log.Error().Msgf("get data candle fail :%s", err.Error())

				continue
			}
		}
	}

	err = s.startEcsServices(validRecord)
	if err != nil {
		return fmt.Errorf("start ecs fail: %w", err)
	}

	return nil
}

//startEcsServices start ecs services
func (s *service) startEcsServices(records []model.Record) error {
	for i, record := range records {
		if record.Type == model.FirstTypeRunning {
			err := s.executeFirstTypeRunning(&records[i])
			if err != nil {
				return fmt.Errorf("first type running error : %w", err)
			}

			return nil
		}

		err := s.executeSecondTypeRunning(&records[i])
		if err != nil {
			return fmt.Errorf("seconds type running error : %w", err)
		}
	}

	return nil
}

//executeFirstTypeRunning start bp03 with 2 input params
func (s *service) executeFirstTypeRunning(data *model.Record) error {
	command := "./" + s.Cfg.ShellPath
	_, err := exec.Command(command, data.Type, data.Kubun, data.Hasshin, data.CreateDay, data.CreateTime, data.StartIndex).Output() //nolint:gosec
	if err != nil {
		return fmt.Errorf("error execute first type running %w", err)
	}

	return nil
}

//executeSecondTypeRunning start bp03 with 2 input params
func (s *service) executeSecondTypeRunning(data *model.Record) error {
	command := "./" + s.Cfg.ShellPath
	_, err := exec.Command(command, data.Type, data.PathFolder).Output() //nolint:gosec
	if err != nil {
		return err
	}

	return nil
}
