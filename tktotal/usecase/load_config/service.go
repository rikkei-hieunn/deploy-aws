/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"tktotal/configs"
	"tktotal/infrastructure/repository"
	"tktotal/model"
)

//Service structure load Config
type Service struct {
	Config       *configs.Server
	S3Repository repository.IS3Repository
}

// NewService service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) IConfigurationLoader {
	return &Service{
		Config:       cfg,
		S3Repository: s3Repo,
	}
}

// LoadConfig load Config process
func (s *Service) LoadConfig(ctx context.Context, date string) error {
	var err error
	dates, err := s.LoadConfigDate(date)
	if err != nil {
		return err
	}
	s.Config.Dates = dates
	suybetus, err := s.ParseSyubetu()
	if err != nil {
		return err
	}
	s.Config.Suybetu = append(s.Config.Suybetu, suybetus...)

	return nil
}

//LoadConfigDate getting date
func (s *Service) LoadConfigDate(date string) ([]time.Time, error) {
	//get specific day
	var result []time.Time
	if date != model.EmptyString {
		specificDate, err := time.Parse(model.DateFormatWithoutStroke, date)
		if err != nil {
			return nil, fmt.Errorf("invalid params date : %w", err)
		}
		result = append(result, specificDate)

		return result, nil
	}
	//get 7 days previous current day
	today := time.Now()
	for i := 6; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		result = append(result, date)
	}

	return result, nil
}

//ParseSyubetu parse syubetu
func (s *Service) ParseSyubetu() ([]string, error) {
	var err error
	type res struct {
		Suybetu []string `json:"Syubetu"`
	}
	suybetu := res{}
	dataByte, err := os.ReadFile(s.Config.SyubetuFileDefinition)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataByte, &suybetu)
	if err != nil {
		return nil, err
	}

	return suybetu.Suybetu, err
}
