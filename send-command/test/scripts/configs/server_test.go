package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"send-command/configs"
	"testing"
)

// Init
func Test_Init(t *testing.T) {
	type args struct {
		path     string
		fileName string
	}
	type expect struct {
		cfg *configs.Server
		err error
	}

	tests := []struct {
		name string
		args
		expect
	}{
		{
			name: "parse environment file success",
			args: args{
				path:     ".",
				fileName: "environment_variables.json",
			},
			expect: expect{
				cfg: &configs.Server{
					TickSystem: configs.TickSystem{
						GroupsDefinitionKei1Object: "environment_variables/groups_definition_first_kei.json",
						GroupsDefinitionKei2Object: "environment_variables/groups_definition_second_kei.json",
						DevelopEnvironment:         true,
					},
				},
				err: nil,
			},
		},
		{
			name: "parse environment file failed due to file not found",
			args: args{
				path:     ".",
				fileName: "environment_variables1.json",
			},
			expect: expect{
				cfg: nil,
				err: viper.ConfigFileNotFoundError{},
			},
		},
		{
			name: "parse environment file failed due to wrong format",
			args: args{
				path:     ".",
				fileName: "environment_variables_wrong_format.json",
			},
			expect: expect{
				cfg: nil,
				err: &mapstructure.Error{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := configs.Init(tt.args.path, tt.args.fileName)

			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", tt.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, tt.expect.cfg.GroupsDefinitionKei1Object, result.GroupsDefinitionKei1Object)
				assert.Equal(t, tt.expect.cfg.GroupsDefinitionKei2Object, result.GroupsDefinitionKei2Object)
				assert.Equal(t, tt.expect.cfg.DevelopEnvironment, result.DevelopEnvironment)
			}
		})
	}
}
