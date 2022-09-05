/*
Package configs contains configuration.
*/
package configs

import (
	"fmt"
	"recreate-one-minute/model"
)

// TickSystem struct contain info application
type TickSystem struct {
	CommonDefinitionObject      string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	DB1EndpointDefinitionObject string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	DB2EndpointDefinitionObject string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	DevelopEnvironment          bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`
	ShellPath                   string `mapstructure:"TK_SYSTEM_START_EC2_SHELL_PATH"`
	Kei                         string
	TickRegion                  string
	TickBucket                  string
	CandleTablePrefix           string
}

// OneMinuteCommon structure define prefix of table candle management
type OneMinuteCommon struct {
	CandleManagementTablePrefix string `json:"TK_CANDLE_MANAGEMENT_TABLE_NAME"`
}

// Validate validate log config
func (s *TickSystem) Validate() error {
	if s.CommonDefinitionObject == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}
	if s.DB1EndpointDefinitionObject == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required")
	}
	if s.DB2EndpointDefinitionObject == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required")
	}
	if s.ShellPath == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_START_EC2_SHELL_PATH required")
	}

	return nil
}
