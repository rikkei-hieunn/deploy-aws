package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	QuoteCodesDefinitionTickKei1Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionKehaiKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	QuoteCodesDefinitionKehaiKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT"`
	DatabaseStatusDefinitionObject      string `mapstructure:"TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT"`
	DevelopEnvironment                  bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	Kei         string
	Request     interface{}
	RequestType string
	DataType    string
	S3Region    string
	S3Bucket    string
}

// QuoteCodes quote code from environment variable file
type QuoteCodes struct {
	QKbn     string `json:"QKBN"`
	Sndc     string `json:"SNDC"`
	LogicID  string `json:"TKLOGIC_ID"`
	Endpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName   string `json:"TKDBNAME"`
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
	if len(c.QuoteCodesDefinitionTickKei1Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required")
	}
	if len(c.QuoteCodesDefinitionKehaiKei1Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT required")
	}
	if len(c.QuoteCodesDefinitionTickKei2Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required")
	}
	if len(c.QuoteCodesDefinitionKehaiKei2Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT required")
	}
	if len(c.DatabaseStatusDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT required")
	}

	return nil
}
