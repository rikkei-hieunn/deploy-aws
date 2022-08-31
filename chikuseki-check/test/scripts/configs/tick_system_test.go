package configs

import (
	"chikuseki-check/configs"
	"chikuseki-check/model"
	"errors"
	"github.com/stretchr/testify/assert"
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
				QuoteCodesDefinitionTickKei1Object:  model.EmptyString,
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				CommonDefinitionObject:              "environment_variables/common_variables.json",
				DevelopEnvironment:                  true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: model.EmptyString,
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				CommonDefinitionObject:              "environment_variables/common_variables.json",
				DevelopEnvironment:                  true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  model.EmptyString,
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				CommonDefinitionObject:              "environment_variables/common_variables.json",
				DevelopEnvironment:                  true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: model.EmptyString,
				CommonDefinitionObject:              "environment_variables/common_variables.json",
				DevelopEnvironment:                  true,
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_SHARE_INFORMATION_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/elements_definition.json",
				CommonDefinitionObject:              model.EmptyString,
				DevelopEnvironment:                  true,
			},
			expect: errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/elements_definition.json",
				CommonDefinitionObject:              "environment_variables/common_variables.json",
				DevelopEnvironment:                  true,
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}
}
