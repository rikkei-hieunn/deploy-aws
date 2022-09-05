package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"recreate-one-minute/configs"
	"testing"
)

func TestInit(t *testing.T) {

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
			}{path: ".", name: "environment_variables.json"},
			expect: &configs.Server{
				TickSystem: configs.TickSystem{
					CommonDefinitionObject:      "environment_variables/common_variables.json",
					DB1EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei1.json",
					DB2EndpointDefinitionObject: "environment_variables/quote_code_definition_tick_kei2.json",
					DevelopEnvironment:          true,
					ShellPath:                   "/home/ec2-user/BP-08/start-ecs",
				},
			},
			err: nil,
		}, {
			name: "parse environment variables file error wrong data format",
			args: struct {
				path string
				name string
			}{path: ".", name: "environment_variables_wrong_format.json"},
			expect: nil,
			err:    &mapstructure.Error{},
		}, {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := configs.Init(tt.args.path, tt.args.name)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, result.TickSystem.CommonDefinitionObject, tt.expect.CommonDefinitionObject)
				assert.Equal(t, result.TickSystem.DB1EndpointDefinitionObject, tt.expect.DB1EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.DB2EndpointDefinitionObject, tt.expect.DB2EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
				assert.Equal(t, result.TickSystem.ShellPath, tt.expect.ShellPath)
			}
		})
	}
}
