package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"update-status/configs"
)

func Test_Tick_System_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "validate success",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				DatabaseStatusDefinitionObject:      "environment_variables/database_status_definition.json",
			},
			expect: nil,
		},
		{
			name: "missing TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				DatabaseStatusDefinitionObject:      "environment_variables/database_status_definition.json",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required"),
		},
		{
			name: "missing TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				DatabaseStatusDefinitionObject:      "environment_variables/database_status_definition.json",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT required"),
		},
		{
			name: "missing TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				DatabaseStatusDefinitionObject:      "environment_variables/database_status_definition.json",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required"),
		},
		{
			name: "missing TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_kehai_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "",
				DatabaseStatusDefinitionObject:      "environment_variables/database_status_definition.json",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT required"),
		},
		{
			name: "missing TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_kehai_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				DatabaseStatusDefinitionObject:      "",
			},
			expect: errors.New("system TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.args.Validate()
			assert.Equal(t, result, test.expect)
		})
	}
}
