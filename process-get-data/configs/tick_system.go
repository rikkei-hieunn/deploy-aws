package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	PathNewCronTabs                    string `mapstructure:"TK_SYSTEM_PATH_NEW_CRON_TABS"`
	PathStartBP03                      string `mapstructure:"TK_SYSTEM_PATH_START_BP03"`
	QuoteCodesDefinitionTickKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	OneMinuteOperatorConfigObject      string `mapstructure:"TK_SYSTEM_ONE_MINUTE_OPERATOR_CONFIG_DEFINITION_OBJECT"`
	CommonDefinitionObject             string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	DevelopEnvironment                 bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region       string
	S3Bucket       string
	ProcessGetData string
}

// OneMinuteCommon structure define prefix of table candle management
type OneMinuteCommon struct {
	CandleManagementTablePrefix string `json:"TK_CANDLE_MANAGEMENT_TABLE_NAME"`
}

// OneMinuteConfig structure define information for one minute config
type OneMinuteConfig struct {
	QKbn             string `json:"QKBN"`
	Sndc             string `json:"SNDC"`
	OperatorType     string `json:"OPERATION_TYPE"`
	OriginStartIndex string
	StartIndex       string `json:"START_INDEX"`
	EndIndex         string `json:"END_INDEX"`
	QuoteCode        string `json:"QCD"`
	CreateTime       string `json:"CREATE_TIME"`
	CreateDay        int    `json:"CREATE_DAY"`
	StartTime        string `json:"START_TIME"`
	EndTime          string `json:"END_TIME"`
	TableName        string `json:"TABLE_NAME"`
	Mon              string `json:"MON"`
	Tue              string `json:"TUE"`
	Wed              string `json:"WED"`
	Thu              string `json:"THU"`
	Fri              string `json:"FRI"`
	Sat              string `json:"SAT"`
	Sun              string `json:"SUN"`
}

// CronInfo info create cron tab
type CronInfo struct {
	Minute      string
	Hour        string
	DayOfMonth  int
	Month       int
	DayOfWeek   int
	Command     string
	RequestType int
	StartIndex  string
	QKBN        string
	SNDC        string
	CreateDate  string
	CreateTime  string
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
	QKbn           string `json:"QKBN"`
	Sndc           string `json:"SNDC"`
	LogicID        string `json:"TKLOGIC_ID"`
	Endpoint       string `json:"TKDB_MASTER_ENDPOINT"`
	DBName         string `json:"TKDBNAME"`
	NumberRange    Range  `json:"NUMBER_RANGES"`
	CharacterRange Range  `json:"CHARACTER_RANGES"`
}

// Range range definition
type Range struct {
	Start string `json:"TKSTART_PROPERNAME"`
	End   string `json:"TKEND_PROPERNAME"`
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
	TkQkbn      string    `json:"TKQKBN"`
	TablePrefix string    `json:"TKPREFIX"`
	Elements    []Element `json:"TKELEMENTS"`
	StartDate   string    `json:"TKSTART_DATE"`
	EndDate     string    `json:"TKEND_DATE"`
}

// Validate validate config
func (c *TickSystem) Validate() error {
	if len(c.QuoteCodesDefinitionTickKei1Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required")
	}
	if len(c.QuoteCodesDefinitionTickKei2Object) == 0 {
		return errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required")
	}
	if len(c.CommonDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}
	if len(c.OneMinuteOperatorConfigObject) == 0 {
		return errors.New("system TK_SYSTEM_ONE_MINUTE_OPERATOR_CONFIG_DEFINITION_OBJECT required")
	}

	return nil
}
