package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	GroupsDefinitionObject              string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_OBJECT"`
	QuoteCodesDefinitionTickKei1Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionKehaiKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	QuoteCodesDefinitionKehaiKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT"`
	ElementsDefinitionObject            string `mapstructure:"TK_SYSTEM_ELEMENTS_DEFINITION_OBJECT"`
	CommonDefinitionObject              string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	OneMinuteDefinitionObject           string `mapstructure:"TK_SYSTEM_ONE_MINUTE_DEFINITION_OBJECT"`
	DevelopEnvironment                  bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region   string
	S3Bucket   string
	CreateType string
}

// OneMinuteCommon structure define prefix of table one minute
type OneMinuteCommon struct {
	OneMinuteTablePrefix string `json:"TK_ONE_MINUTE_DATA_TABLE_NAME"`
}

// TargetCreateTable structure define information for create a new table
type TargetCreateTable struct {
	LogicGroup  string
	QKbn        string
	Sndc        string
	Endpoint    string
	DBName      string
	DataType    string
	TablePrefix string
}

// Group group definition
type Group struct {
	LogicGroup  string `json:"TKLOGIC_GROUP"`
	TopicName   string `json:"TKTOPIC_NAME"`
	Types       string `json:"TKTYPES"`
	TypesString []string
}

// QuoteCodes quote code from environment variable file
type QuoteCodes struct {
	QKbn     string `json:"QKBN"`
	Sndc     string `json:"SNDC"`
	LogicID  string `json:"TKLOGIC_ID"`
	Endpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName   string `json:"TKDBNAME"`
}

// Element element definition
type Element struct {
	Name   string `mapstructure:"ELEMENT" json:"ELEMENT"`
	Column string `mapstructure:"COLUMN" json:"COLUMN"`
	Length int    `mapstructure:"LENGTH" json:"LENGTH"`
}

// TableDefinition table definition
type TableDefinition struct {
	Elements  []Element `json:"TKELEMENTS"`
	StartDate string    `json:"TKSTART_DATE"`
	EndDate   string    `json:"TKEND_DATE"`
}

// OneMinuteTableDefinition structure define table one minute
type OneMinuteTableDefinition struct {
	TkQkbn    string    `json:"TKQKBN"`
	Elements  []Element `json:"TKELEMENTS"`
	StartDate string    `json:"TKSTART_DATE"`
	EndDate   string    `json:"TKEND_DATE"`
}

// Validate validate config
func (c *TickSystem) Validate() error {
	if len(c.GroupsDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_GROUP_DEFINITION_OBJECT required")
	}
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
	if len(c.ElementsDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_ELEMENTS_DEFINITION_OBJECT required")
	}
	if len(c.CommonDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}
	if len(c.OneMinuteDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_ONE_MINUTE_DEFINITION_OBJECT required")
	}

	return nil
}
