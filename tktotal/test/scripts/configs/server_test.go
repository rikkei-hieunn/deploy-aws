package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"tktotal/configs"
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
					InfoLogPath:           "rikkei-work/20220815/BP-04_Logs/Info/",
					ErrorLogPath:          "rikkei-work/20220815/BP-04_Logs/Error/",
					OutputLogPath:         "rikkei-work/20220815/BP-04_Logs/Output/",
					SyubetuFileDefinition: "environment_variables/suybetu_information_definition.json",
					Port: []string{
						"9000",
						"9001",
						"9002",
					},
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
				assert.Equal(t, result.TickSystem.InfoLogPath, tt.expect.InfoLogPath)
				assert.Equal(t, result.TickSystem.ErrorLogPath, tt.expect.ErrorLogPath)
				assert.Equal(t, result.TickSystem.OutputLogPath, tt.expect.OutputLogPath)
				assert.Equal(t, result.TickSystem.SyubetuFileDefinition, tt.expect.SyubetuFileDefinition)
				assert.Equal(t, len(result.TickSystem.Port), len(tt.expect.Port))
				for index := range result.TickSystem.Port {
					assert.Equal(t, result.TickSystem.Port[index], tt.expect.Port[index])
				}
			}
		})
	}
}
