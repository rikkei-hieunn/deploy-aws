package configs

import (
	"chikuseki-check/configs"
	"chikuseki-check/model"
	"errors"
	"github.com/stretchr/testify/assert"
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
