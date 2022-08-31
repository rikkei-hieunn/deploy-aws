package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"message-receive-check/configs"
	"testing"
)

func Test_Init(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			path string
			name string
		}
		expect *configs.Server
		err    error
	}{
		{
			name: "parse environment variables file success",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables.json",
			},
			expect: &configs.Server{
				TickSystem: configs.TickSystem{
					GroupsDefinitionObject:    "environment_variables/groups_definition.json",
					MessageCountLogKei1Object: "environment_variables/kei1",
					MessageCountLogKei2Object: "environment_variables/kei2",
					NumberPercentAlert:        20,
					DevelopEnvironment:        true,
				},
			},
			err: nil,
		},
		{
			name: "parse environment variables file error file not found",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables1.json",
			},
			expect: nil,
			err:    viper.ConfigFileNotFoundError{},
		},
		{
			name: "parse environment variables file error wrong data format",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables_wrong_format.json",
			},
			expect: nil,
			err:    &mapstructure.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := configs.Init(tt.args.path, tt.args.name)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, result.TickSystem.GroupsDefinitionObject, tt.expect.GroupsDefinitionObject)
				assert.Equal(t, result.TickSystem.MessageCountLogKei1Object, tt.expect.MessageCountLogKei1Object)
				assert.Equal(t, result.TickSystem.MessageCountLogKei2Object, tt.expect.MessageCountLogKei2Object)
				assert.Equal(t, result.TickSystem.NumberPercentAlert, tt.expect.NumberPercentAlert)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
			}
		})
	}
}
