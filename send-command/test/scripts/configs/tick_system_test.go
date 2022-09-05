package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"send-command/configs"
	"testing"
)

// Validate
func Test_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "validate success",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: "environment_variables/groups_definition_first_kei.json",
				GroupsDefinitionKei2Object: "environment_variables/groups_definition_second_kei.json",
				DevelopEnvironment:         true,
			},
			expect: nil,
		},
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: "",
				GroupsDefinitionKei2Object: "environment_variables/groups_definition_second_kei.json",
				DevelopEnvironment:         true,
			},
			expect: errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: "environment_variables/groups_definition_first_kei.json",
				GroupsDefinitionKei2Object: "",
				DevelopEnvironment:         true,
			},
			expect: errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.args.Validate())
		})
	}
}
