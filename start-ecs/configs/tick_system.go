/*
Package configs implements configuration for program
*/
package configs

import (
	"errors"
	"fmt"
)

//TickSystem contain config for program
type TickSystem struct {
	S3Region string
	S3Bucket string

	QuoteCodesDefinitionTickKei1Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT"`
	QuoteCodesDefinitionKehaiKei1Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT"`
	QuoteCodesDefinitionTickKei2Object  string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT"`
	QuoteCodesDefinitionKehaiKei2Object string `mapstructure:"TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT"`

	BP03FirstRunningTypeEnvKeys  []string `mapstructure:"TK_SYSTEM_BP03_PARAMS_FIRST_RUNNING_TYPE_ENV_KEYS"`
	BP03SecondRunningTypeEnvKeys []string `mapstructure:"TK_SYSTEM_BP03_PARAMS_SECOND_RUNNING_TYPE_ENV_KEYS"`

	BP05FirstRunningTypeEnvKeys  []string `mapstructure:"TK_SYSTEM_BP05_PARAMS_FIRST_RUNNING_ENV_KEYS"`
	BP05SecondRunningTypeEnvKeys []string `mapstructure:"TK_SYSTEM_BP05_PARAMS_SECOND_RUNNING_ENV_KEYS"`
	BP05ThirdRunningTypeEnvKeys  []string `mapstructure:"TK_SYSTEM_BP05_PARAMS_THIRD_RUNNING_ENV_KEYS"`
	BP05FourthRunningTypeEnvKeys []string `mapstructure:"TK_SYSTEM_BP05_PARAMS_FOURTH_RUNNING_ENV_KEYS"`

	BP06EnvKeyParams []string `mapstructure:"TK_SYSTEM_BP06_PARAMS_ENV_KEYS"`

	BP07FirstRunningTypeEnvKeys  []string `mapstructure:"TK_SYSTEM_BP07_PARAMS_FIRST_RUNNING_TYPE_ENV_KEYS"`
	BP07SecondRunningTypeEnvKeys []string `mapstructure:"TK_SYSTEM_BP07_PARAMS_SECOND_RUNNING_TYPE_ENV_KEYS"`
	BP07ThirdRunningTypeEnvKeys  []string `mapstructure:"TK_SYSTEM_BP07_PARAMS_THIRD_RUNNING_TYPE_ENV_KEYS"`

	DevelopEnvironment bool `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	MaxWaitTime   int `mapstructure:"TK_SYSTEM_MAX_RETRY_WAIT_TIME"`
	RetryWaitTime int `mapstructure:"TK_SYSTEM_RETRY_WAIT_TIME"`
	MaxCountTime  int `mapstructure:"TK_SYSTEM_MAX_CHECK_COUNT_TIME"`
}

// QuoteCodes quote code from environment variable file
type QuoteCodes struct {
	QKbn     string `json:"QKBN"`
	Sndc     string `json:"SNDC"`
	LogicID  string `json:"TKLOGIC_ID"`
	Endpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName   string `json:"TKDBNAME"`
}

//Validate validate configuration
func (t *TickSystem) Validate() error {
	if len(t.QuoteCodesDefinitionTickKei1Object) == 0 {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is required")
	}
	if len(t.QuoteCodesDefinitionKehaiKei1Object) == 0 {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT is required")
	}
	if len(t.QuoteCodesDefinitionTickKei2Object) == 0 {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is required")
	}
	if len(t.QuoteCodesDefinitionKehaiKei2Object) == 0 {
		return errors.New("TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT is required")
	}
	if len(t.BP03FirstRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP03_PARAMS_FIRST_RUNNING_TYPE_ENV_KEYS is required")
	}
	if len(t.BP03SecondRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP03_PARAMS_SECOND_RUNNING_TYPE_ENV_KEYS is required")
	}
	if len(t.BP05FirstRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP05_PARAMS_FIRST_RUNNING_ENV_KEYS is required")
	}
	if len(t.BP05SecondRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP05_PARAMS_SECOND_RUNNING_ENV_KEYS is required")
	}
	if len(t.BP05ThirdRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP05_PARAMS_THIRD_RUNNING_ENV_KEYS is required")
	}
	if len(t.BP05FourthRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP05_PARAMS_FOURTH_RUNNING_ENV_KEYS is required")
	}
	if len(t.BP06EnvKeyParams) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP06_PARAMS_ENV_KEYS is required")
	}

	if len(t.BP07FirstRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP07_PARAMS_FIRST_RUNNING_TYPE_ENV_KEYS is required")
	}
	if len(t.BP07SecondRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP07_PARAMS_SECOND_RUNNING_TYPE_ENV_KEYS is required")
	}
	if len(t.BP07ThirdRunningTypeEnvKeys) == 0 {
		return fmt.Errorf("TK_SYSTEM_BP07_PARAMS_THIRD_RUNNING_TYPE_ENV_KEYS is required")
	}
	if t.MaxWaitTime == 0 {
		return fmt.Errorf("TK_SYSTEM_MAX_RETRY_WAIT_TIME is required")
	}
	if t.RetryWaitTime == 0 {
		return fmt.Errorf("TK_SYSTEM_RETRY_WAIT_TIME is required")
	}
	if t.MaxCountTime == 0 {
		return fmt.Errorf("TK_SYSTEM_MAX_CHECK_COUNT_TIME is required")
	}

	return nil
}
