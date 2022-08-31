package configs

import (
	"data-del/configs"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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
					CommonDefinitionObject:           "environment_variables/common_variables.json",
					TickDB1EndpointDefinitionObject:  "environment_variables/quote_code_definition_tick_first_kei.json",
					TickDB2EndpointDefinitionObject:  "environment_variables/quote_code_definition_tick_second_kei.json",
					KehaiDB1EndpointDefinitionObject: "environment_variables/quote_code_definition_kehai_first_kei.json",
					KehaiDB2EndpointDefinitionObject: "environment_variables/quote_code_definition_kehai_second_kei.json",
					ExpireDefinitionObject:           "environment_variables/expired_definition.json",
					DevelopEnvironment:               true,
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
				assert.Equal(t, result.TickSystem.CommonDefinitionObject, tt.expect.CommonDefinitionObject)
				assert.Equal(t, result.TickSystem.ExpireDefinitionObject, tt.expect.ExpireDefinitionObject)
				assert.Equal(t, result.TickSystem.TickDB1EndpointDefinitionObject, tt.expect.TickDB1EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.TickDB2EndpointDefinitionObject, tt.expect.TickDB2EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.KehaiDB1EndpointDefinitionObject, tt.expect.KehaiDB1EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.KehaiDB2EndpointDefinitionObject, tt.expect.KehaiDB2EndpointDefinitionObject)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
			}
		})
	}
}
