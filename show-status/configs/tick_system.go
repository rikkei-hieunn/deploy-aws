package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	DatabaseStatusDefinitionObject string `mapstructure:"TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT"`
	DevelopEnvironment             bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region string
	S3Bucket string
}

// DatabaseStatus define structure contain status of two kei database
type DatabaseStatus struct {
	QKbn               string `json:"QKBN"`
	Sndc               string `json:"SNDC"`
	TheFirstKeiStatus  bool   `json:"TK_STATUS_KEI1"`
	TheSecondKeiStatus bool   `json:"TK_STATUS_KEI2"`
}

// ArrayDatabaseStatus define list of database status
type ArrayDatabaseStatus []DatabaseStatus

// GroupDatabaseStatusDefinition structure define group database status
type GroupDatabaseStatusDefinition struct {
	Tick  ArrayDatabaseStatus `json:"tick_data"`
	Kehai ArrayDatabaseStatus `json:"kehai_data"`
}

// Validate validate object tick system
func (c *TickSystem) Validate() error {
	if len(c.DatabaseStatusDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT required")
	}

	return nil
}
