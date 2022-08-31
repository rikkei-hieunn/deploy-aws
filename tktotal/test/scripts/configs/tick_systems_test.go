package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tktotal/configs"
)

func Test_TickSystem_Validate(t *testing.T) {
	var tests = []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "Validate success",
			args: configs.TickSystem{
				InfoLogPath:           "infor-log",
				ErrorLogPath:          "error path",
				OutputLogPath:         "output log path",
				Port:                  []string{"111", "4444"},
				SyubetuFileDefinition: "ssssssssssssss",
			},
			expect: nil,
		},
		{
			name: "TK_SYSTEM_TOIAWASE_INFO_LOG_PATH is empty",
			args: configs.TickSystem{
				InfoLogPath:           "",
				ErrorLogPath:          "error path",
				OutputLogPath:         "output log path",
				Port:                  []string{"111", "4444"},
				SyubetuFileDefinition: "ssssssssssssss",
			},
			expect: fmt.Errorf("TK_SYSTEM_TOIAWASE_INFO_LOG_PATH is required "),
		},
		{
			name: "TK_SYSTEM_TOIAWASE_ERROR_LOG_PATH is empty",
			args: configs.TickSystem{
				InfoLogPath:           "aaaa",
				ErrorLogPath:          "",
				OutputLogPath:         "output log path",
				Port:                  []string{"111", "4444"},
				SyubetuFileDefinition: "ssssssssssssss",
			},
			expect: fmt.Errorf("TK_SYSTEM_TOIAWASE_ERROR_LOG_PATH is required "),
		},
		{
			name: "TK_SYSTEM_TOIAWASE_OUTPUT_LOG_PATH is empty",
			args: configs.TickSystem{
				InfoLogPath:           "aaaa",
				ErrorLogPath:          "aaaa",
				OutputLogPath:         "",
				Port:                  []string{"111", "4444"},
				SyubetuFileDefinition: "ssssssssssssss",
			},
			expect: fmt.Errorf("TK_SYSTEM_TOIAWASE_OUTPUT_LOG_PATH is required "),
		},
		{
			name: "TK_SYSTEM_PORT is empty",
			args: configs.TickSystem{
				InfoLogPath:           "aaaa",
				ErrorLogPath:          "aaaa",
				OutputLogPath:         "xxxxxxx",
				Port:                  []string{},
				SyubetuFileDefinition: "ssssssssssssss",
			},
			expect: fmt.Errorf("TK_SYSTEM_PORT is required "),
		},
		{
			name: "TK_SYSTEM_SYUBETU_INFORMATION_DEFINITION is empty",
			args: configs.TickSystem{
				InfoLogPath:           "aaaa",
				ErrorLogPath:          "aaaa",
				OutputLogPath:         "xxxxxxx",
				Port:                  []string{"a"},
				SyubetuFileDefinition: "",
			},
			expect: fmt.Errorf("TK_SYSTEM_SYUBETU_INFORMATION_DEFINITION is required "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}
}
