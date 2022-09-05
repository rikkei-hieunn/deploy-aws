/*
Package configs contains configuration.
*/
package configs

import "fmt"

// TickSystem structure of config about system
type TickSystem struct {
	CommonDefinitionObject  string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	CalendarFileName1       string `mapstructure:"TK_SYSTEM_CALENDAR1_FILE_NAME"`
	CalendarFileName2       string `mapstructure:"TK_SYSTEM_CALENDAR2_FILE_NAME"`
	FilebusDefinitionObject string `mapstructure:"TK_SYSTEM_FILEBUS_DEFINITION_OBJECT"`
	DevelopEnvironment      bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region          string
	S3Bucket          string
	CalendarTableName string
}

// Validate validate Tick System config
func (c *TickSystem) Validate() error {
	if len(c.CommonDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}
	if len(c.CalendarFileName1) == 0 {
		return fmt.Errorf("system TK_SYSTEM_CALENDAR1_FILE_NAME required")
	}
	if len(c.CalendarFileName2) == 0 {
		return fmt.Errorf("system TK_SYSTEM_CALENDAR2_FILE_NAME required")
	}
	if len(c.FilebusDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_FILEBUS_DEFINITION_OBJECT required")
	}

	return nil
}
