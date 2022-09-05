package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/model"
	"testing"
)

func Test_TickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_CALENDAR1_FILE_NAME is empty",
			args: configs.TickSystem{
				CalendarFileName1:       model.EmptyString,
				CalendarFileName2:       "calendar_file2.csv",
				CommonDefinitionObject:  "environment_variables/common_variables.json",
				FilebusDefinitionObject: "environment_variables/filebus_definition.json",
				DevelopEnvironment:      true,
			},
			expect: errors.New("system TK_SYSTEM_CALENDAR1_FILE_NAME required"),
		},
		{
			name: "TK_SYSTEM_CALENDAR2_FILE_NAME is empty",
			args: configs.TickSystem{
				CalendarFileName1:       "calendar_file1.csv",
				CalendarFileName2:       model.EmptyString,
				CommonDefinitionObject:  "environment_variables/common_variables.json",
				FilebusDefinitionObject: "environment_variables/filebus_definition.json",
				DevelopEnvironment:      true,
			},
			expect: errors.New("system TK_SYSTEM_CALENDAR2_FILE_NAME required"),
		},
		{
			name: "TK_SYSTEM_SHARE_INFORMATION_OBJECT is empty",
			args: configs.TickSystem{
				CalendarFileName1:       "calendar_file1.csv",
				CalendarFileName2:       "calendar_file2.csv",
				CommonDefinitionObject:  model.EmptyString,
				FilebusDefinitionObject: "environment_variables/filebus_definition.json",
				DevelopEnvironment:      true,
			},
			expect: errors.New("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required"),
		},
		{
			name: "TK_SYSTEM_FILEBUS_DEFINITION_OBJECT is empty",
			args: configs.TickSystem{
				CalendarFileName1:       "calendar_file1.csv",
				CalendarFileName2:       "calendar_file2.csv",
				CommonDefinitionObject:  "environment_variables/common_variables.json",
				FilebusDefinitionObject: model.EmptyString,
				DevelopEnvironment:      true,
			},
			expect: errors.New("system TK_SYSTEM_FILEBUS_DEFINITION_OBJECT required"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				CalendarFileName1:       "calendar_file1.csv",
				CalendarFileName2:       "calendar_file2.csv",
				CommonDefinitionObject:  "environment_variables/common_variables.json",
				FilebusDefinitionObject: "environment_variables/filebus_definition.json",
				DevelopEnvironment:      true,
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}
}
