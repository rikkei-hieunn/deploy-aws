package configs

import (
	"data-del/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Validate_tick_db(t *testing.T) {
	var tests = []struct {
		name   string
		args   configs.TickDB
		expect string
	}{
		{
			name: "port missing",
			args: configs.TickDB{
				Port:              0,
				User:              "tick",
				Password:          "test",
				MaxOpenConnection: 1,
				MaxIdleConnection: 1,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "database TK_DB_PORT required",
		},
		{
			name: "user missing",
			args: configs.TickDB{
				Port:              3456,
				User:              "",
				Password:          "test",
				MaxOpenConnection: 1,
				MaxIdleConnection: 1,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "database TK_DB_USER required",
		},
		{
			name: "password missing",
			args: configs.TickDB{
				Port:              3456,
				User:              "tick",
				Password:          "",
				MaxOpenConnection: 1,
				MaxIdleConnection: 1,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "database TK_DB_PASSWORD required",
		},
		{
			name: "MaxOpenConnection missing",
			args: configs.TickDB{
				Port:              3456,
				User:              "tick",
				Password:          "test",
				MaxOpenConnection: 0,
				MaxIdleConnection: 1,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "database TK_DB_MAX_OPEN_CONNECTION required",
		},
		{
			name: "MaxIdleConnection missing",
			args: configs.TickDB{
				Port:              3456,
				User:              "tick",
				Password:          "test",
				MaxOpenConnection: 1,
				MaxIdleConnection: 0,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "database TK_DB_MAX_IDLE_CONNECTION required",
		},
		{
			name: "DriverName missing",
			args: configs.TickDB{
				Port:              3456,
				User:              "tick",
				Password:          "test",
				MaxOpenConnection: 1,
				MaxIdleConnection: 1,
				DriverName:        "",
				RetryTimes:        1,
			},
			expect: "database TK_DB_DRIVER_NAME required",
		},
		{
			name: "validate success",
			args: configs.TickDB{
				Port:              3456,
				User:              "tick",
				Password:          "test",
				MaxOpenConnection: 1,
				MaxIdleConnection: 1,
				DriverName:        "mysql",
				RetryTimes:        1,
			},
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			if tt.name != "validate success" {
				assert.Equal(t, result.Error(), tt.expect)
			} else {
				assert.Equal(t, result, nil)
			}

		})
	}
}
