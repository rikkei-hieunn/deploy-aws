/*
Package startinstance implements logics repository.
*/
package startinstance

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"start-jushin/infrastructure/repository"
	"start-jushin/model"
)

type service struct {
	ssm    repository.ISSMRepository
}

//NewService aws ssm constructor
func NewService(ssm repository.ISSMRepository) IExecuteProgram {
	return &service{ssm: ssm}
}

//ExecuteProgram start execute ssm command
func (s *service) ExecuteProgram(ctx context.Context) error {
	result, err := s.ssm.ExecuteProgram(ctx)
	if result == nil {
		return fmt.Errorf("process fails.... ")
	}
	var isFail bool
	if err != nil {
		return err
	}
	for insID, result := range result {
		if result == model.EmptyString {
			isFail = true
			log.Error().Msgf("Start instance ID %s fail ", insID)
		}
	}
	if isFail {
		return fmt.Errorf("process fails.... ")
	}

	return nil
}
