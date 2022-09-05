package configs

import (
	"errors"
	"fmt"
	"start-jushin/model"
	"strings"
)

// TickSystem contains application configuration
type TickSystem struct {
	GroupDefinitionKei1Object           string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT"`
	GroupDefinitionKei2Object           string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT"`
	QuoteCodesDefinitionTickKei1Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionKehaiKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	QuoteCodesDefinitionKehaiKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT"`
	InstancePathKey                     string `mapstructure:"TK_SYSTEM_INSTANCE_PATH_ENV_KEY"`
	DevelopEnvironment                  bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	Kei      string
	S3Bucket string
	S3Region string
}

// QuoteCodes quote code from environment variable file
type QuoteCodes struct {
	QKbn     string `json:"QKBN"`
	Sndc     string `json:"SNDC"`
	LogicID  string `json:"TKLOGIC_ID"`
	Endpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName   string `json:"TKDBNAME"`
}

// Validate func validates application configuration
func (s *TickSystem) Validate() error {
	if strings.TrimSpace(s.QuoteCodesDefinitionTickKei1Object) == model.EmptyString {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is required")
	}
	if strings.TrimSpace(s.QuoteCodesDefinitionKehaiKei1Object) == model.EmptyString {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT is required")
	}
	if strings.TrimSpace(s.QuoteCodesDefinitionTickKei2Object) == model.EmptyString {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is required")
	}
	if strings.TrimSpace(s.QuoteCodesDefinitionKehaiKei2Object) == model.EmptyString {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT is required")
	}
	if strings.TrimSpace(s.GroupDefinitionKei1Object) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT is required")
	}
	if strings.TrimSpace(s.GroupDefinitionKei2Object) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT is required")
	}
	if strings.TrimSpace(s.InstancePathKey) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_INSTANCE_PATH_ENV_KEY is required")
	}

	return nil
}
