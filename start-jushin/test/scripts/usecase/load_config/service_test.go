package load_config

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"start-jushin/configs"
	loadconfig "start-jushin/usecase/load_config"
	"testing"
)

func initService() *loadconfig.Service {
	tickConfig := new(configs.Server)
	tickConfig.DevelopEnvironment = true
	service := new(loadconfig.Service)
	service.Configs = tickConfig

	return service
}

func Test_LoadGroupData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			configDB []configs.Group
			err      error
		}
	}{
		{
			name: "group definition file not found",
			path: ".",
			expect: struct {
				configDB []configs.Group
				err      error
			}{
				configDB: nil,
				err:      new(fs.PathError),
			},
		},
		{
			name: "group definition file not found",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				configDB []configs.Group
				err      error
			}{
				configDB: nil,
				err:      new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse group definition file successfully",
			path: "definition_files/groups_definition.json",
			expect: struct {
				configDB []configs.Group
				err      error
			}{
				configDB: []configs.Group{
					{
						GroupID:       "TE1",
						TopicName:     "SZN-TSE1",
						Types:         "1,2,4",
						CommandPort:   "7000",
						TickHostName:  "ATCKDR1",
						KehaiHostName: "ATCKDR61",
						GroupLine:     "TSE",
					},
					{
						GroupID:       "TE2",
						TopicName:     "SZN-TSE2",
						Types:         "1,2,4",
						CommandPort:   "7001",
						TickHostName:  "ATCKDR12",
						KehaiHostName: "ATCKDR61",
						GroupLine:     "TSE",
					},
					{
						GroupID:       "XXX",
						TopicName:     "SZN-TSE2",
						Types:         "1,2,4",
						CommandPort:   "7002",
						TickHostName:  "ATCKDR3",
						KehaiHostName: "ATCKDR61",
						GroupLine:     "TSE",
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := service.LoadGroupData(test.path)
			if err != nil {
				assert.Equal(t, test.expect.configDB, result)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.err, err)
				assert.Equal(t, len(test.expect.configDB), len(result))
				for i := range result {
					assert.Equal(t, test.expect.configDB[i].TopicName, result[i].TopicName)
					assert.Equal(t, test.expect.configDB[i].GroupID, result[i].GroupID)
					assert.Equal(t, test.expect.configDB[i].Types, result[i].Types)
					assert.Equal(t, test.expect.configDB[i].CommandPort, result[i].CommandPort)
					assert.Equal(t, test.expect.configDB[i].TickHostName, result[i].TickHostName)
					assert.Equal(t, test.expect.configDB[i].KehaiHostName, result[i].KehaiHostName)
					assert.Equal(t, test.expect.configDB[i].GroupLine, result[i].GroupLine)
				}
			}
		})
	}
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
