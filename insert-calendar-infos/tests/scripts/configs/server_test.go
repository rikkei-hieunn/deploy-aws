package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/configs"
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
					CalendarFileName1:       "calendar_file1.csv",
					CalendarFileName2:       "calendar_file2.csv",
					CommonDefinitionObject:  "environment_variables/common_variables.json",
					FilebusDefinitionObject: "environment_variables/filebus_definition.json",
					DevelopEnvironment:      true,
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
				assert.Equal(t, result.TickSystem.CalendarFileName1, tt.expect.CalendarFileName1)
				assert.Equal(t, result.TickSystem.CalendarFileName2, tt.expect.CalendarFileName2)
				assert.Equal(t, result.TickSystem.CommonDefinitionObject, tt.expect.CommonDefinitionObject)
				assert.Equal(t, result.TickSystem.FilebusDefinitionObject, tt.expect.FilebusDefinitionObject)
				assert.Equal(t, result.TickSystem.DevelopEnvironment, tt.expect.DevelopEnvironment)
			}
		})
	}
}
