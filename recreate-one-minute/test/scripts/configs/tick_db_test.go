package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"recreate-one-minute/configs"
	"testing"
)

func TestTickDB_Validate(t *testing.T) {

	tests := []struct {
		name   string
		args   configs.TickDB
		expect error
	}{
		{
			name: "validate successfully",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
			},
			expect: nil,
		},
		{
			name: "TK_DB_PORT is missing",
			args: configs.TickDB{
				Port:              0,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
			},
			expect: errors.New("database TK_DB_PORT required"),
		},
		{
			name: "TK_DB_USER is missing",
			args: configs.TickDB{
				Port:              3306,
				User:              "",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
			},
			expect: errors.New("database TK_DB_USER required"),
		},
		{
			name: "TK_DB_PASSWORD is missing",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
			},
			expect: errors.New("database TK_DB_PASSWORD required"),
		},
		{
			name: "TK_DB_MAX_OPEN_CONNECTION is missing",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 0,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
			},
			expect: errors.New("database TK_DB_MAX_OPEN_CONNECTION required"),
		},
		{
			name: "TK_DB_MAX_IDLE_CONNECTION is missing",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 0,
				DriverName:        "mysql",
			},
			expect: errors.New("database TK_DB_MAX_IDLE_CONNECTION required"),
		},
		{
			name: "TK_DB_DRIVER_NAME is missing",
			args: configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "",
			},
			expect: errors.New("database TK_DB_DRIVER_NAME required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}

}
