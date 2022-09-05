package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"process-get-data/configs"
	"process-get-data/model"
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
			expect: errors.New("TK_DB_PORT"),
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
			expect: errors.New("TK_DB_USER"),
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
			expect: errors.New("TK_DB_PASSWORD"),
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
			expect: errors.New("TK_DB_MAX_OPEN_CONNECTION"),
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
			expect: errors.New("TK_DB_MAX_IDLE_CONNECTION"),
		},
		{
			name: "TK_DB_DRIVER_NAME is empty",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "",
				RetryTimes:        3,
				RetryWaitMs:       10000,
			},
			expect: errors.New("TK_DB_DRIVER_NAME"),
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
