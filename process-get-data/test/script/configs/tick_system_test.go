package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"process-get-data/configs"
	"testing"
)

func Test_TickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object: "",
				QuoteCodesDefinitionTickKei2Object: "environment_variables/quote_code_definition_tick_kei2.json",
				CommonDefinitionObject:             "environment_variables/common_variables.json",
				OneMinuteOperatorConfigObject:      "one_minute_operation_configs.json",
				DevelopEnvironment:                 true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object: "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionTickKei2Object: "",
				CommonDefinitionObject:             "environment_variables/common_variables.json",
				OneMinuteOperatorConfigObject:      "one_minute_operation_configs.json",
				DevelopEnvironment:                 true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_SHARE_INFORMATION_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object: "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionTickKei2Object: "environment_variables/quote_code_definition_tick_kei2.json",
				CommonDefinitionObject:             "",
				OneMinuteOperatorConfigObject:      "one_minute_operation_configs.json",
				DevelopEnvironment:                 true,
			},
			expect: errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_ONE_MINUTE_OPERATOR_CONFIG_DEFINITION_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object: "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionTickKei2Object: "environment_variables/quote_code_definition_tick_kei2.json",
				CommonDefinitionObject:             "environment_variables/common_variables.json",
				OneMinuteOperatorConfigObject:      "",
				DevelopEnvironment:                 true,
			},
			expect: errors.New("system TK_SYSTEM_ONE_MINUTE_OPERATOR_CONFIG_DEFINITION_OBJECT required"),
		},
		{
			name: "Validate successfully",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object: "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionTickKei2Object: "environment_variables/quote_code_definition_tick_kei2.json",
				CommonDefinitionObject:             "environment_variables/common_variables.json",
				OneMinuteOperatorConfigObject:      "one_minute_operation_configs.json",
				DevelopEnvironment:                 true,
			},
			expect: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.args.Validate()
			assert.Equal(t, result, test.expect)
			//if result != nil {
			//
			//}
		})
	}
}
