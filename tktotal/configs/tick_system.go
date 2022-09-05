package configs

import (
	"fmt"
	"time"
	"tktotal/model"
)

//TickSystem tick systems configuration
type TickSystem struct {
	InfoLogPath           string   `mapstructure:"TK_SYSTEM_TOIAWASE_INFO_LOG_PATH"`
	ErrorLogPath          string   `mapstructure:"TK_SYSTEM_TOIAWASE_ERROR_LOG_PATH"`
	OutputLogPath         string   `mapstructure:"TK_SYSTEM_TOIAWASE_OUTPUT_LOG_PATH"`
	Port                  []string `mapstructure:"TK_SYSTEM_PORT"`
	SyubetuFileDefinition string   `mapstructure:"TK_SYSTEM_SYUBETU_INFORMATION_DEFINITION"`

	Suybetu  []string `json:"Syubetu"`
	Dates    []time.Time
	S3Region string
	S3Bucket string
}

//Validate validate configuration
func (s *TickSystem) Validate() error {
	if s.InfoLogPath == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_TOIAWASE_INFO_LOG_PATH is required ")
	}
	if s.ErrorLogPath == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_TOIAWASE_ERROR_LOG_PATH is required ")
	}
	if s.OutputLogPath == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_TOIAWASE_OUTPUT_LOG_PATH is required ")
	}
	if s.SyubetuFileDefinition == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_SYUBETU_INFORMATION_DEFINITION is required ")
	}
	if len(s.Port) == 0 {
		return fmt.Errorf("TK_SYSTEM_PORT is required ")
	}

	return nil
}
