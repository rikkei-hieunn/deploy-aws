package load_config

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"start-ecs/configs"
	loadconfig "start-ecs/usecase/load_config"
	"testing"
)

func initService() *loadconfig.Service {
	tickConfig := new(configs.Server)
	tickConfig.DevelopEnvironment = true
	service := new(loadconfig.Service)
	service.Configs = tickConfig

	return service
}

func Test_LoadQuoteCodeData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			quoteCodes map[string][]configs.QuoteCodes
			err        error
		}
	}{
		{
			name: "quote code definition file not found",
			path: ".",
			expect: struct {
				quoteCodes map[string][]configs.QuoteCodes
				err        error
			}{
				quoteCodes: nil,
				err:        new(fs.PathError),
			},
		},
		{
			name: "parse quote code definition file wrong format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				quoteCodes map[string][]configs.QuoteCodes
				err        error
			}{
				quoteCodes: nil,
				err:        new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse quote code definition file successfully",
			path: "definition_files/quote_code_definition_kehai_kei1.json",
			expect: struct {
				quoteCodes map[string][]configs.QuoteCodes
				err        error
			}{
				quoteCodes: map[string][]configs.QuoteCodes{
					"CX2": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX20",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
					"CX3": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX30",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := service.LoadQuoteCodeData(test.path)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, len(test.expect.quoteCodes), len(result))
				for key, quoteCodeActual := range result {
					quoteCodeExpect := test.expect.quoteCodes[key]
					assert.Equal(t, len(quoteCodeActual), len(quoteCodeExpect))
					for i := range quoteCodeActual {
						assert.Equal(t, quoteCodeActual[i].QKbn, quoteCodeExpect[i].QKbn)
						assert.Equal(t, quoteCodeActual[i].Sndc, quoteCodeExpect[i].Sndc)
						assert.Equal(t, quoteCodeActual[i].LogicID, quoteCodeExpect[i].LogicID)
						assert.Equal(t, quoteCodeActual[i].Endpoint, quoteCodeExpect[i].Endpoint)
						assert.Equal(t, quoteCodeActual[i].DBName, quoteCodeExpect[i].DBName)
					}
				}
			}
		})
	}
}

func Test_ValidateUniqueEndpoint(t *testing.T) {
	var tests = []struct {
		name       string
		quoteCodes map[string][]configs.QuoteCodes
		expect     bool
	}{
		{
			name:       "Quote Codes Invalid - Wrong Endpoint",
			quoteCodes: createInvalidQCDsWrongEndpoint(),
			expect:     false,
		},
		{
			name:       "Quote Codes Invalid - Wrong DB Name",
			quoteCodes: createInvalidQCDsWrongDBName(),
			expect:     false,
		},
		{
			name:       "Quote Codes Valid",
			quoteCodes: createValidQCDs(),
			expect:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := loadconfig.ValidateUniqueEndpoint(test.quoteCodes)
			assert.Equal(t, result, test.expect)
		})
	}
}

func createValidQCDs() map[string][]configs.QuoteCodes {
	return map[string][]configs.QuoteCodes{
		"CXJ0": {
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ0",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ0",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
	}
}

func createInvalidQCDsWrongEndpoint() map[string][]configs.QuoteCodes {
	return map[string][]configs.QuoteCodes{
		"CXJ1": {
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ1",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ1",
				Endpoint: "127.0.0.2",
				DBName:   "tick",
			},
		},
	}
}

func createInvalidQCDsWrongDBName() map[string][]configs.QuoteCodes {
	return map[string][]configs.QuoteCodes{
		"CXJ1": {
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ1",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
			{
				QKbn:     "E",
				Sndc:     "CXJ",
				LogicID:  "CXJ1",
				Endpoint: "127.0.0.1",
				DBName:   "tick2",
			},
		},
	}
}
