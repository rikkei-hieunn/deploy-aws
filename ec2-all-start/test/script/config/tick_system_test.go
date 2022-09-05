package configs

import (
	"ec2-all-start/configs"
	"ec2-all-start/model"
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
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: model.EmptyString,
				GroupsDefinitionKei2Object: "environment_variables/groups_definition_second_kei.json",
				DevelopEnvironment:         true,
			},
			expect: errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: "environment_variables/groups_definition_first_kei.json",
				GroupsDefinitionKei2Object: model.EmptyString,
				DevelopEnvironment:         true,
			},
			expect: errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT required"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				GroupsDefinitionKei1Object: "environment_variables/groups_definition_first_kei.json",
				GroupsDefinitionKei2Object: "environment_variables/groups_definition_second_kei.json",
				DevelopEnvironment:         true,
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