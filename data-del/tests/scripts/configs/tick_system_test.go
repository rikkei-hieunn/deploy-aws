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
				CommonDefinitionObject:           "configuration_files/common_variables.json",
				ExpireDefinitionObject:           "aaaaa",
				TickDB1EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke1.json",
				KehaiDB1EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke1.json",
				TickDB2EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke2.json",
				KehaiDB2EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke2.json",
			},
			expect: "",
		},
		{
			name: "miss TK_SYSTEM_SHARE_INFORMATION_OBJECT",
			args: configs.TickSystem{
				CommonDefinitionObject:           "",
				ExpireDefinitionObject:           "aaaaa",
				TickDB1EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke1.json",
				KehaiDB1EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke1.json",
				TickDB2EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke2.json",
				KehaiDB2EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke2.json",
			},
			expect: "system TK_SYSTEM_SHARE_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_TICK_KE1_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:           "aaaaa",
				CommonDefinitionObject:           "configuration_files/common_variables.json",
				TickDB1EndpointDefinitionObject:  "",
				KehaiDB1EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke1.json",
				TickDB2EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke2.json",
				KehaiDB2EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke2.json",
			},
			expect: "system TK_SYSTEM_TICK_DB1_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_KEHAI_KE1_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:           "aaaaa",
				CommonDefinitionObject:           "configuration_files/common_variables.json",
				TickDB1EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke1.json",
				KehaiDB1EndpointDefinitionObject: "",
				TickDB2EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke2.json",
				KehaiDB2EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke2.json",
			},
			expect: "system TK_SYSTEM_KEHAI_DB1_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_TICK_KE2_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:           "aaaaa",
				CommonDefinitionObject:           "configuration_files/common_variables.json",
				TickDB1EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke1.json",
				KehaiDB1EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke1.json",
				TickDB2EndpointDefinitionObject:  "",
				KehaiDB2EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke2.json",
			},
			expect: "system TK_SYSTEM_TICK_DB2_ENDPOINT_INFORMATION_OBJECT required",
		},
		{
			name: "miss TK_SYSTEM_QUOTE_CODES_KEHAI_KE2_DEFINITION_OBJECT",
			args: configs.TickSystem{
				ExpireDefinitionObject:           "aaaaa",
				CommonDefinitionObject:           "configuration_files/common_variables.json",
				TickDB1EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke1.json",
				KehaiDB1EndpointDefinitionObject: "configuration_files/quote_code_definition_kehai_ke1.json",
				TickDB2EndpointDefinitionObject:  "configuration_files/quote_code_definition_tick_ke2.json",
				KehaiDB2EndpointDefinitionObject: "",
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
