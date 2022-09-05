package configs

import (
	"data-del/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Validate_tck_sys(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect string
	}{
		{
			name: "Validate success",
			args: configs.TickSystem{
				CommonDefinitionObject:             "configuration_files/common_variables.json",
				ExpireDefinitionObject:             "aaaaa",
				TickKei1QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei1.json",
				KehaiKei1QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei1.json",
				TickKei2QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei2.json",
				KehaiKei2QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei2.json",
			},
			expect: "",
		},
		{
			name: "miss TK_SYSTEM_SHARE_INFORMATION_OBJECT",
			args: configs.TickSystem{
				CommonDefinitionObject:             "",
				ExpireDefinitionObject:             "aaaaa",
				TickKei1QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei1.json",
				KehaiKei1QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei1.json",
				TickKei2QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei2.json",
				KehaiKei2QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei2.json",
			},
			expect: "system TK_SYSTEM_SHARE_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_TICK_KE1_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:             "aaaaa",
				CommonDefinitionObject:             "configuration_files/common_variables.json",
				TickKei1QuoteCodeDefinitionObject:  "",
				KehaiKei1QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei1.json",
				TickKei2QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei2.json",
				KehaiKei2QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei2.json",
			},
			expect: "system TK_SYSTEM_TICK_DB1_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_KEHAI_KE1_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:             "aaaaa",
				CommonDefinitionObject:             "configuration_files/common_variables.json",
				TickKei1QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei1.json",
				KehaiKei1QuoteCodeDefinitionObject: "",
				TickKei2QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei2.json",
				KehaiKei2QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei2.json",
			},
			expect: "system TK_SYSTEM_KEHAI_DB1_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_TICK_KE2_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:             "aaaaa",
				CommonDefinitionObject:             "configuration_files/common_variables.json",
				TickKei1QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei1.json",
				KehaiKei1QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei1.json",
				TickKei2QuoteCodeDefinitionObject:  "",
				KehaiKei2QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei2.json",
			},
			expect: "system TK_SYSTEM_TICK_DB2_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_KEHAI_KE2_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:             "aaaaa",
				CommonDefinitionObject:             "configuration_files/common_variables.json",
				TickKei1QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei1.json",
				KehaiKei1QuoteCodeDefinitionObject: "configuration_files/quote_code_definition_kehai_kei1.json",
				TickKei2QuoteCodeDefinitionObject:  "configuration_files/quote_code_definition_tick_kei2.json",
				KehaiKei2QuoteCodeDefinitionObject: "",
			},
			expect: "system TK_SYSTEM_KEHAI_DB2_ENDPOINT_INFORMATION_OBJECT required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			if tt.name != "Validate success" {
				assert.Equal(t, result.Error(), tt.expect)
			} else {
				assert.Equal(t, result, nil)
			}
		})
	}
}
