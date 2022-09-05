package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	QuoteCodesDefinitionTickKei1Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionKehaiKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	QuoteCodesDefinitionKehaiKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT"`
	CommonDefinitionObject              string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	DevelopEnvironment                  bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region           string
	S3Bucket           string
	Kubun, Hassin, ZXD string
}

// QuoteCodes quote code from environment variable file
type QuoteCodes struct {
	QKbn     string `json:"QKBN"`
	Sndc     string `json:"SNDC"`
	LogicID  string `json:"TKLOGIC_ID"`
	Endpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName   string `json:"TKDBNAME"`
}

// Validate validate config
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
	if len(c.CommonDefinitionObject) == 0 {
		return errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}

	return nil
}
