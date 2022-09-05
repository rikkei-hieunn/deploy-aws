package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"recreate-one-minute/configs"
	"testing"
)

func TestTickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_SHARE_INFORMATION_OBJECT is missing",
			args: configs.TickSystem{
				CommonDefinitionObject:      "",
				DB1EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei1.json",
				DB2EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei2.json",
				ShellPath:                   "/home/ec2-user/BP-08/start-ecs",
			},
			expect: errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required"),
		}, {
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is missing",
			args: configs.TickSystem{
				CommonDefinitionObject:      "environment_variables/common_variables.json",
				DB1EndpointDefinitionObject: "",
				DB2EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei2.json",
				ShellPath:                   "/home/ec2-user/BP-08/start-ecs",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT required"),
		}, {
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is missing",
			args: configs.TickSystem{
				CommonDefinitionObject:      "environment_variables/common_variables.json",
				DB1EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei1.json",
				DB2EndpointDefinitionObject: "",
				ShellPath:                   "/home/ec2-user/BP-08/start-ecs",
			},
			expect: errors.New("system TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT required"),
		}, {
			name: "TK_SYSTEM_START_EC2_SHELL_PATH is missing",
			args: configs.TickSystem{
				CommonDefinitionObject:      "environment_variables/common_variables.json",
				DB1EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei1.json",
				DB2EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei2.json",
				ShellPath:                   "",
			},
			expect: errors.New("system TK_SYSTEM_START_EC2_SHELL_PATH required"),
		}, {
			name: "validate successfully",
			args: configs.TickSystem{
				CommonDefinitionObject:      "environment_variables/common_variables.json",
				DB1EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei1.json",
				DB2EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei2.json",
				ShellPath:                   "/home/ec2-user/BP-08/start-ecs",
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
