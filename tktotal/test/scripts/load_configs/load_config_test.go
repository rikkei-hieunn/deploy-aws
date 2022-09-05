package load_configs

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
	"time"
	"tktotal/configs"
	loadconfig "tktotal/usecase/load_config"
)

func Test_LoadConfigDate(t *testing.T) {
	sv := loadconfig.Service{}
	sv.Config = &configs.Server{
		TickSystem: configs.TickSystem{
			SyubetuFileDefinition: "test.json",
		},
	}
	today := time.Now()
	var mock []string
	for i := 6; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		mock = append(mock, date.Format("20060102"))
	}

	var tests = []struct {
		name   string
		args   string
		expect interface{}
	}{
		{
			name:   "input empty string",
			args:   "",
			expect: mock,
		},
		{
			name:   "invalid date",
			args:   "aaaaa",
			expect: fmt.Errorf("invalid params date : parsing time \"aaaaa\" as \"20060102\": cannot parse \"aaaaa\" as \"2006\""),
		},
		{
			name:   "input is a date",
			args:   "20220808",
			expect: []string{"20220808"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := sv.LoadConfigDate(tt.args)
			if tt.name == "invalid date" {
				assert.EqualValues(t, err.Error(), tt.expect.(error).Error())
			} else {
				expect := tt.expect.([]string)
				for i, re := range res {
					assert.Equal(t, re.Format("20060102"), expect[i])
				}

			}
		})
	}
}
func Test_ParseSyubetu(t *testing.T) {
	sv := loadconfig.Service{}
	sv.Config = &configs.Server{
		TickSystem: configs.TickSystem{
			SyubetuFileDefinition: "test.json",
		},
	}
	var tests = []struct {
		name   string
		args   string
		path   string
		expect []string
		err    error
	}{
		{
			name: "Parse syubetu success",
			path: "definition_files/test.json",
			expect: []string{
				"総件数",
				"1000",
				"1010",
				"102",
				"1020",
				"1100",
			},
			err: nil,
		},
		{
			name:   "Parse syubetu file not found",
			path:   ".",
			expect: nil,
			err:    &fs.PathError{},
		},
		{
			name:   "Parse syubetu wrong data format",
			path:   "definition_files/wrong_format_file.json",
			expect: nil,
			err:    &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv.Config.SyubetuFileDefinition = tt.path
			suybetu, err := sv.ParseSyubetu()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				for i, s := range suybetu {
					assert.Equal(t, s, tt.expect[i])
				}
			}
		})
	}
}
