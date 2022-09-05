package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"start-jushin/configs"
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
					QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
					GroupDefinitionKei1Object:           "environment_variables/groups_definition_first_kei.json",
					GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
					InstancePathKey:                     "TK_JUSHIN_PROCESS_PATH",
					DevelopEnvironment:                  true,
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
				assert.Equal(t, result.TickSystem.GroupDefinitionKei1Object, tt.expect.GroupDefinitionKei1Object)
				assert.Equal(t, result.TickSystem.GroupDefinitionKei2Object, tt.expect.GroupDefinitionKei2Object)
				assert.Equal(t, result.TickSystem.InstancePathKey, tt.expect.InstancePathKey)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
			}
		})
	}
}