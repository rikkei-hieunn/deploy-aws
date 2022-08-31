package load_config

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"message-receive-check/configs"
	"message-receive-check/model"
	loadconfig "message-receive-check/usecase/load_config"
	"strings"
	"testing"
)

func initService() *loadconfig.Service {
	tickConfig := new(configs.Server)
	tickConfig.DevelopEnvironment = true
	service := new(loadconfig.Service)
	service.Config = tickConfig

	return service
}

func Test_LoadGroupData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			processNames []string
			groups       []configs.Group
			err          error
		}
	}{
		{
			name: "group definition file not found",
			path: ".",
			expect: struct {
				processNames []string
				groups       []configs.Group
				err          error
			}{
				processNames: nil,
				groups:       nil,
				err:          new(fs.PathError),
			},
		},
		{
			name: "group definition wrong data format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				processNames []string
				groups       []configs.Group
				err          error
			}{
				processNames: nil,
				groups:       nil,
				err:          new(json.UnmarshalTypeError),
			},
		},
		{
			name: "group definition empty data",
			path: "definition_files/groups_definition_empty_data.json",
			expect: struct {
				processNames []string
				groups       []configs.Group
				err          error
			}{
				processNames: nil,
				groups:       nil,
				err:          nil,
			},
		},
		{
			name: "group definition valid data",
			path: "definition_files/groups_definition.json",
			expect: struct {
				processNames []string
				groups       []configs.Group
				err          error
			}{
				processNames: []string{"TE1", "KTE1", "TE2", "KTE2"},
				groups: []configs.Group{
					{
						LogicGroup: "TE1",
						Types:      "1,2,4",
					},
					{
						LogicGroup: "TE2",
						Types:      "1,2,4",
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.GroupsDefinitionObject = test.path
			processNames, groups, err := service.LoadGroupData()
			if err != nil {
				assert.Equal(t, test.expect.processNames, processNames)
				assert.Equal(t, test.expect.groups, groups)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.err, err)
				assert.Equal(t, len(test.expect.processNames), len(processNames))
				allNames := model.CommaCharacter + strings.Join(test.expect.processNames, model.CommaCharacter) + model.CommaCharacter
				for i := range processNames {
					assert.Contains(t, allNames, model.CommaCharacter+processNames[i]+model.CommaCharacter)
				}
				assert.Equal(t, len(test.expect.groups), len(groups))
				for i := range groups {
					assert.Equal(t, test.expect.groups[i].LogicGroup, groups[i].LogicGroup)
					assert.Equal(t, test.expect.groups[i].Types, groups[i].Types)
				}
			}
		})
	}
}
