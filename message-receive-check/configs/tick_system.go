package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	GroupsDefinitionObject    string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_OBJECT"`
	MessageCountLogKei1Object string `mapstructure:"TK_SYSTEM_MESSAGE_COUNT_LOG_KEI1_PATH"`
	MessageCountLogKei2Object string `mapstructure:"TK_SYSTEM_MESSAGE_COUNT_LOG_KEI2_PATH"`
	NumberPercentAlert        int    `mapstructure:"TK_SYSTEM_NUMBER_PERCENT_ALERT"`
	DevelopEnvironment        bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region string
	S3Bucket string
}

// Group definition
type Group struct {
	LogicGroup string `json:"TKLOGIC_GROUP"`
	Types      string `json:"TKTYPES"`
}

// Validate validate config
func (c *TickSystem) Validate() error {
	if c.NumberPercentAlert == 0 {
		return errors.New("system TK_SYSTEM_NUMBER_PERCENT_ALERT required")
	}
	if c.NumberPercentAlert < 0 {
		return errors.New("invalid TK_SYSTEM_NUMBER_PERCENT_ALERT")
	}
	if len(c.GroupsDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_GROUP_DEFINITION_OBJECT required")
	}
	if len(c.MessageCountLogKei1Object) == 0 {
		return errors.New("system TK_SYSTEM_MESSAGE_COUNT_LOG_KEI1_PATH required")
	}
	if len(c.MessageCountLogKei2Object) == 0 {
		return errors.New("system TK_SYSTEM_MESSAGE_COUNT_LOG_KEI2_PATH required")
	}

	return nil
}
