package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"show-status/configs"
	"testing"
)

func Test_Tick_System_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT is empty",
			args: configs.TickSystem{
				DatabaseStatusDefinitionObject: "",
				DevelopEnvironment:             true,
			},
			expect: errors.New("system TK_SYSTEM_DATABASE_STATUS_DEFINITION_OBJECT required"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				DatabaseStatusDefinitionObject: "environment_variables/database_status_definition.json",
				DevelopEnvironment:             true,
			},
			expect: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.args.Validate()
			assert.Equal(t, result, test.expect)
		})
	}
}
