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
					CommonDefinitionObject:             "environment_variables/common_variables.json",
					TickKei1QuoteCodeDefinitionObject:  "environment_variables/quote_code_definition_tick_kei1.json",
					TickKei2QuoteCodeDefinitionObject:  "environment_variables/quote_code_definition_tick_kei2.json",
					KehaiKei1QuoteCodeDefinitionObject: "environment_variables/quote_code_definition_kehai_kei1.json",
					KehaiKei2QuoteCodeDefinitionObject: "environment_variables/quote_code_definition_kehai_kei2.json",
					ExpireDefinitionObject:             "environment_variables/expired_definition.json",
					DevelopEnvironment:                 true,
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
				assert.Equal(t, result.TickSystem.TickKei1QuoteCodeDefinitionObject, tt.expect.TickKei1QuoteCodeDefinitionObject)
				assert.Equal(t, result.TickSystem.TickKei2QuoteCodeDefinitionObject, tt.expect.TickKei2QuoteCodeDefinitionObject)
				assert.Equal(t, result.TickSystem.KehaiKei1QuoteCodeDefinitionObject, tt.expect.KehaiKei1QuoteCodeDefinitionObject)
				assert.Equal(t, result.TickSystem.KehaiKei2QuoteCodeDefinitionObject, tt.expect.KehaiKei2QuoteCodeDefinitionObject)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
			}
		})
	}
}
