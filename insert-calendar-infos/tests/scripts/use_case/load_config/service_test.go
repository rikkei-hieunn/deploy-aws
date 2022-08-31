package load_config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/configs"
	loadconfig "insert-calendar-infos/usecase/load_config"
	"io/fs"
	"testing"
)

func initService() *loadconfig.Service {
	tickConfig := new(configs.Server)
	tickConfig.DevelopEnvironment = true
	service := new(loadconfig.Service)
	service.Config = tickConfig

	return service
}

func Test_LoadCommonData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			configDB   *configs.TickDB
			calendarDB *configs.CalendarInfo
			err        error
		}
	}{
		{
			name: "common definition file not found",
			path: ".",
			expect: struct {
				configDB   *configs.TickDB
				calendarDB *configs.CalendarInfo
				err        error
			}{
				configDB:   nil,
				calendarDB: nil,
				err:        new(fs.PathError),
			},
		},
		{
			name: "common definition wrong file format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				configDB   *configs.TickDB
				calendarDB *configs.CalendarInfo
				err        error
			}{
				configDB:   nil,
				calendarDB: nil,
				err:        new(json.UnmarshalTypeError),
			},
		},
		{
			name: "common definition validate error",
			path: "definition_files/common_variables_error.json",
			expect: struct {
				configDB   *configs.TickDB
				calendarDB *configs.CalendarInfo
				err        error
			}{
				configDB:   nil,
				calendarDB: nil,
				err:        errors.New("database TK_DB_PORT required"),
			},
		},
		{
			name: "load common definition file successfully",
			path: "definition_files/common_variables.json",
			expect: struct {
				configDB   *configs.TickDB
				calendarDB *configs.CalendarInfo
				err        error
			}{
				configDB: &configs.TickDB{
					Port:              3306,
					User:              "admin",
					Password:          "123456123",
					MaxOpenConnection: 100,
					MaxIdleConnection: 100,
					DriverName:        "mysql",
					RetryTimes:        3,
					RetryWaitMs:       3000,
				},
				calendarDB: &configs.CalendarInfo{
					TableName:      "calendar_infos",
					DBKei1Endpoint: "10.1.40.148",
					DBKei1Name:     "tick",
					DBKei2Endpoint: "10.1.40.178",
					DBKei2Name:     "tick",
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.CommonDefinitionObject = test.path
			tickDB, calendarInfo, err := service.LoadCommonData()
			if err != nil {
				assert.Equal(t, test.expect.configDB, tickDB)
				assert.Equal(t, test.expect.calendarDB, calendarInfo)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.configDB.Port, tickDB.Port)
				assert.Equal(t, test.expect.configDB.User, tickDB.User)
				assert.Equal(t, test.expect.configDB.Password, tickDB.Password)
				assert.Equal(t, test.expect.configDB.MaxOpenConnection, tickDB.MaxOpenConnection)
				assert.Equal(t, test.expect.configDB.MaxIdleConnection, tickDB.MaxIdleConnection)
				assert.Equal(t, test.expect.configDB.DriverName, tickDB.DriverName)
				assert.Equal(t, test.expect.configDB.RetryTimes, tickDB.RetryTimes)
				assert.Equal(t, test.expect.configDB.RetryWaitMs, tickDB.RetryWaitMs)
				assert.Equal(t, test.expect.calendarDB.TableName, calendarInfo.TableName)
				assert.Equal(t, test.expect.calendarDB.DBKei1Endpoint, calendarInfo.DBKei1Endpoint)
				assert.Equal(t, test.expect.calendarDB.DBKei1Name, calendarInfo.DBKei1Name)
				assert.Equal(t, test.expect.calendarDB.DBKei2Endpoint, calendarInfo.DBKei2Endpoint)
				assert.Equal(t, test.expect.calendarDB.DBKei2Name, calendarInfo.DBKei2Name)
			}
		})
	}
}

func Test_LoadFilebusData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			filebus *configs.TickFileBus
			err     error
		}
	}{
		{
			name: "filebus definition file not found",
			path: ".",
			expect: struct {
				filebus *configs.TickFileBus
				err     error
			}{
				filebus: nil,
				err:     new(fs.PathError),
			},
		},
		{
			name: "filebus definition wrong file format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				filebus *configs.TickFileBus
				err     error
			}{
				filebus: nil,
				err:     new(json.UnmarshalTypeError),
			},
		},
		{
			name: "common definition validate error",
			path: "definition_files/filebus_definition_error.json",
			expect: struct {
				filebus *configs.TickFileBus
				err     error
			}{
				filebus: nil,
				err:     errors.New("database TK_FILEBUS_HOST required"),
			},
		},
		{
			name: "load common definition file successfully",
			path: "definition_files/filebus_definition.json",
			expect: struct {
				filebus *configs.TickFileBus
				err     error
			}{
				filebus: &configs.TickFileBus{
					Port:           8089,
					Username:       "admin",
					Password:       "admin",
					Hostname:       "10.1.40.178",
					URLDownload:    "/file/download",
					PathCalendar1:  "sample-file/",
					PathCalendar2:  "sample-file/",
					RetryTimes:     3,
					RetryWaitTimes: 3000,
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.FilebusDefinitionObject = test.path
			filebus, err := service.LoadFilebusData()
			if err != nil {
				assert.Equal(t, test.expect.filebus, filebus)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.filebus.Username, filebus.Username)
				assert.Equal(t, test.expect.filebus.Password, filebus.Password)
				assert.Equal(t, test.expect.filebus.Hostname, filebus.Hostname)
				assert.Equal(t, test.expect.filebus.Port, filebus.Port)
				assert.Equal(t, test.expect.filebus.URLDownload, filebus.URLDownload)
				assert.Equal(t, test.expect.filebus.PathCalendar1, filebus.PathCalendar1)
				assert.Equal(t, test.expect.filebus.PathCalendar2, filebus.PathCalendar2)
				assert.Equal(t, test.expect.filebus.RetryTimes, filebus.RetryTimes)
				assert.Equal(t, test.expect.filebus.RetryWaitTimes, filebus.RetryWaitTimes)
			}
		})
	}
}
