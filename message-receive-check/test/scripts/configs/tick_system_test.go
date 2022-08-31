package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"message-receive-check/configs"
	"message-receive-check/model"
	"testing"
)

func Test_TickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_OBJECT is empty",
			args: configs.TickSystem{
				GroupsDefinitionObject:    model.EmptyString,
				MessageCountLogKei1Object: "environment_variables/kei1",
				MessageCountLogKei2Object: "environment_variables/kei1",
				NumberPercentAlert:        20,
				DevelopEnvironment:        true,
			},
			expect: errors.New("system TK_SYSTEM_GROUP_DEFINITION_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_MESSAGE_COUNT_LOG_KEI1_PATH is empty",
			args: configs.TickSystem{
				GroupsDefinitionObject:    "environment_variables/groups_definition.json",
				MessageCountLogKei1Object: model.EmptyString,
				MessageCountLogKei2Object: "environment_variables/kei1",
				NumberPercentAlert:        20,
				DevelopEnvironment:        true,
			},
			expect: errors.New("system TK_SYSTEM_MESSAGE_COUNT_LOG_KEI1_PATH required"),
		},
		{
			name: "TK_SYSTEM_MESSAGE_COUNT_LOG_KEI2_PATH is empty",
			args: configs.TickSystem{
				GroupsDefinitionObject:    "environment_variables/groups_definition.json",
				MessageCountLogKei1Object: "environment_variables/kei1",
				MessageCountLogKei2Object: model.EmptyString,
				NumberPercentAlert:        20,
				DevelopEnvironment:        true,
			},
			expect: errors.New("system TK_SYSTEM_MESSAGE_COUNT_LOG_KEI2_PATH required"),
		},
		{
			name: "TK_SYSTEM_NUMBER_PERCENT_ALERT is empty",
			args: configs.TickSystem{
				GroupsDefinitionObject:    "environment_variables/groups_definition.json",
				MessageCountLogKei1Object: "environment_variables/kei1",
				MessageCountLogKei2Object: "environment_variables/kei1",
				NumberPercentAlert:        0,
				DevelopEnvironment:        true,
			},
			expect: errors.New("system TK_SYSTEM_NUMBER_PERCENT_ALERT required"),
		},
		{
			name: "invalid TK_SYSTEM_NUMBER_PERCENT_ALERT",
			args: configs.TickSystem{
				GroupsDefinitionObject:    "environment_variables/groups_definition.json",
				MessageCountLogKei1Object: "environment_variables/kei1",
				MessageCountLogKei2Object: "environment_variables/kei1",
				NumberPercentAlert:        -10,
				DevelopEnvironment:        true,
			},
			expect: errors.New("invalid TK_SYSTEM_NUMBER_PERCENT_ALERT"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				GroupsDefinitionObject:    "environment_variables/groups_definition.json",
				MessageCountLogKei1Object: "environment_variables/kei1",
				MessageCountLogKei2Object: "environment_variables/kei2",
				NumberPercentAlert:        20,
				DevelopEnvironment:        true,
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
