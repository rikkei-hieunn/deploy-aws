package configs

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"kanshi-process/configs"
	"testing"
)

func Test_Init(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			path string
			name string
		}
		expect *configs.Server
		err    error
	}{
		{
			name: "parse environment variables file success",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables.json",
			},
			expect: &configs.Server{
				TickSystem: configs.TickSystem{
					RequestKinou:     "01",
					RequestKanriID:   "KNR_ID",
					RequestUserID:    "Tick_Test",
					RequestSyubetu:   "2101",
					RequestQuoteCode: "XJPY/4",
					RequestFromDate:  "00000000",
					RequestToDate:    "00000000",
					RequestFunasi:    "1",
					RequestKikan:     "1",
				},
			},
			err: nil,
		},
		{
			name: "parse environment variables file error file not found",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables1.json",
			},
			expect: nil,
			err:    viper.ConfigFileNotFoundError{},
		},
		{
			name: "parse environment variables file error wrong data format",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables_wrong_format.json",
			},
			expect: nil,
			err:    &mapstructure.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := configs.Init(tt.args.path, tt.args.name)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, result.TickSystem.RequestKinou, tt.expect.RequestKinou)
				assert.Equal(t, result.TickSystem.RequestKanriID, tt.expect.RequestKanriID)
				assert.Equal(t, result.TickSystem.RequestUserID, tt.expect.RequestUserID)
				assert.Equal(t, result.TickSystem.RequestSyubetu, tt.expect.RequestSyubetu)
				assert.Equal(t, result.TickSystem.RequestQuoteCode, tt.expect.RequestQuoteCode)
				assert.Equal(t, result.TickSystem.RequestFromDate, tt.expect.RequestFromDate)
				assert.Equal(t, result.TickSystem.RequestToDate, tt.expect.RequestToDate)
				assert.Equal(t, result.TickSystem.RequestFunasi, tt.expect.RequestFunasi)
				assert.Equal(t, result.TickSystem.RequestKikan, tt.expect.RequestKikan)
			}
		})
	}
}
