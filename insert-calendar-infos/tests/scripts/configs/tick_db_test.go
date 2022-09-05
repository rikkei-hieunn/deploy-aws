package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/model"
	"testing"
)

func Test_TickDB_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickDB
		expect error
	}{
		{
			name: "TK_DB_PORT is empty",
			args: configs.TickDB{
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("database TK_DB_PORT required"),
		},
		{
			name: "TK_DB_PORT is negative",
			args: configs.TickDB{
				Port:              -100,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("invalid TK_DB_PORT"),
		},
		{
			name: "TK_DB_USER is empty",
			args: configs.TickDB{
				Port:              3306,
				User:              model.EmptyString,
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("database TK_DB_USER required"),
		},
		{
			name: "TK_DB_PASSWORD is empty",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          model.EmptyString,
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("database TK_DB_PASSWORD required"),
		},
		{
			name: "TK_DB_MAX_OPEN_CONNECTION is empty",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 0,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("database TK_DB_MAX_OPEN_CONNECTION required"),
		},
		{
			name: "TK_DB_MAX_OPEN_CONNECTION is negative",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: -100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("invalid TK_DB_MAX_OPEN_CONNECTION"),
		},
		{
			name: "TK_DB_MAX_IDLE_CONNECTION is empty",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 0,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("database TK_DB_MAX_IDLE_CONNECTION required"),
		},
		{
			name: "TK_DB_MAX_IDLE_CONNECTION is negative",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: -100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("invalid TK_DB_MAX_IDLE_CONNECTION"),
		},
		{
			name: "validate success",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitMs:       10000,
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

func Test_CalendarInfo_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.CalendarInfo
		expect error
	}{
		{
			name: "TK_CALENDAR_TABLE_NAME is empty",
			args: configs.CalendarInfo{
				TableName:      model.EmptyString,
				DBKei1Endpoint: "10.1.40.178",
				DBKei1Name:     "tick",
				DBKei2Endpoint: "10.1.40.178",
				DBKei2Name:     "tick",
			},
			expect: errors.New("database TK_CALENDAR_TABLE_NAME required"),
		},
		{name: "TK_CALENDAR_KEI1_ENDPOINT is empty",
			args: configs.CalendarInfo{
				TableName:      "calendar_infos",
				DBKei1Endpoint: model.EmptyString,
				DBKei1Name:     "tick",
				DBKei2Endpoint: "10.1.40.178",
				DBKei2Name:     "tick",
			},
			expect: errors.New("database TK_CALENDAR_KEI1_ENDPOINT required"),
		},
		{name: "TK_CALENDAR_KEI1_DBNAME is empty",
			args: configs.CalendarInfo{
				TableName:      "calendar_infos",
				DBKei1Endpoint: "10.1.40.178",
				DBKei1Name:     model.EmptyString,
				DBKei2Endpoint: "10.1.40.178",
				DBKei2Name:     "tick",
			},
			expect: errors.New("database TK_CALENDAR_KEI1_DBNAME required"),
		},
		{name: "TK_CALENDAR_KEI2_ENDPOINT is empty",
			args: configs.CalendarInfo{
				TableName:      "calendar_infos",
				DBKei1Endpoint: "10.1.40.178",
				DBKei1Name:     "tick",
				DBKei2Endpoint: model.EmptyString,
				DBKei2Name:     "tick",
			},
			expect: errors.New("database TK_CALENDAR_KEI2_ENDPOINT required"),
		},
		{name: "TK_CALENDAR_KEI2_DBNAME is empty",
			args: configs.CalendarInfo{
				TableName:      "calendar_infos",
				DBKei1Endpoint: "10.1.40.178",
				DBKei1Name:     "tick",
				DBKei2Endpoint: "10.1.40.178",
				DBKei2Name:     model.EmptyString,
			},
			expect: errors.New("database TK_CALENDAR_KEI2_DBNAME required"),
		},
		{name: "validate successfully",
			args: configs.CalendarInfo{
				TableName:      "calendar_infos",
				DBKei1Endpoint: "10.1.40.178",
				DBKei1Name:     "tick",
				DBKei2Endpoint: "10.1.40.178",
				DBKei2Name:     "tick",
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
