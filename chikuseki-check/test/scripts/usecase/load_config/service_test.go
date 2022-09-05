package load_config

import (
	"chikuseki-check/configs"
	"chikuseki-check/model"
	loadconfig "chikuseki-check/usecase/load_config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
)

func initService() *loadconfig.Service {
	tickConfig := new(configs.Server)
	tickConfig.DevelopEnvironment = true
	service := new(loadconfig.Service)
	service.Config = tickConfig

	return service
}

func Test_LoadCommonData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			configDB        *configs.TickDB
			tableNamePrefix *configs.TableNamePrefix
			kubuns          map[string]string
			err             error
		}
	}{
		{
			name: "common definition file not found",
			path: ".",
			expect: struct {
				configDB        *configs.TickDB
				tableNamePrefix *configs.TableNamePrefix
				kubuns          map[string]string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				err:             new(fs.PathError),
			},
		},
		{
			name: "common definition wrong file format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				configDB        *configs.TickDB
				tableNamePrefix *configs.TableNamePrefix
				kubuns          map[string]string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				err:             new(json.UnmarshalTypeError),
			},
		},
		{
			name: "common definition validate error",
			path: "definition_files/common_variables_error.json",
			expect: struct {
				configDB        *configs.TickDB
				tableNamePrefix *configs.TableNamePrefix
				kubuns          map[string]string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				err:             errors.New("database TK_DB_PORT required"),
			},
		},
		{
			name: "load common definition file successfully",
			path: "definition_files/common_variables.json",
			expect: struct {
				configDB        *configs.TickDB
				tableNamePrefix *configs.TableNamePrefix
				kubuns          map[string]string
				err             error
			}{
				configDB: &configs.TickDB{
					Port:              3306,
					User:              "admin",
					Password:          "123456123",
					MaxOpenConnection: 100,
					MaxIdleConnection: 100,
					DriverName:        "mysql",
					RetryTimes:        3,
					RetryWaitMs:       3000,
				},
				tableNamePrefix: &configs.TableNamePrefix{
					{
						DataType: 0,
						Prefix:   "business_day_information_data",
					},
					{
						DataType: 1,
						Prefix:   "best_quote_data",
					},
					{
						DataType: 2,
						Prefix:   "multiple_quote_data",
					},
					{
						DataType: 3,
						Prefix:   "master_quote_data",
					},
					{
						DataType: 4,
						Prefix:   "money_flow_data",
					},
					{
						DataType: 5,
						Prefix:   "option_extended_data",
					},
					{
						DataType: 6,
						Prefix:   "index_trading_data",
					},
				},
				kubuns: map[string]string{
					"@":  "A1",
					"@@": "A2",
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.CommonDefinitionObject = test.path
			tickDB, tablePrefixes, kubunMap, err := service.LoadCommonData()
			if err != nil {
				assert.Equal(t, test.expect.configDB, tickDB)
				assert.Equal(t, test.expect.tableNamePrefix, tablePrefixes)
				assert.Equal(t, test.expect.kubuns, kubunMap)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.configDB.Port, tickDB.Port)
				assert.Equal(t, test.expect.configDB.User, tickDB.User)
				assert.Equal(t, test.expect.configDB.Password, tickDB.Password)
				assert.Equal(t, test.expect.configDB.MaxOpenConnection, tickDB.MaxOpenConnection)
				assert.Equal(t, test.expect.configDB.MaxIdleConnection, tickDB.MaxIdleConnection)
				assert.Equal(t, test.expect.configDB.DriverName, tickDB.DriverName)
				assert.Equal(t, test.expect.configDB.RetryTimes, tickDB.RetryTimes)
				assert.Equal(t, test.expect.configDB.RetryWaitMs, tickDB.RetryWaitMs)
				assert.Equal(t, len(*test.expect.tableNamePrefix), len(*tablePrefixes))
				for i := range *tablePrefixes {
					assert.Equal(t, (*test.expect.tableNamePrefix)[i].Prefix, (*tablePrefixes)[i].Prefix)
				}
				assert.Equal(t, len(test.expect.kubuns), len(kubunMap))
				for key := range kubunMap {
					assert.Equal(t, test.expect.kubuns[key], kubunMap[key])
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

func Test_GetQuoteCodes(t *testing.T) {
	service := initService()

	var tests = []struct {
		name                            string
		kubun, hassin                   string
		tickQuoteCodes, kehaiQuoteCodes map[string][]configs.QuoteCodes
		expect                          map[string]configs.QuoteCodes
	}{
		{
			name:            "Kubun is empty - Return empty map",
			kubun:           model.EmptyString,
			hassin:          "T",
			tickQuoteCodes:  createTickMap(),
			kehaiQuoteCodes: createKehaiMap(),
			expect:          map[string]configs.QuoteCodes{},
		},
		{
			name:            "Hassin is empty - Return empty map",
			kubun:           "E",
			hassin:          model.EmptyString,
			tickQuoteCodes:  createTickMap(),
			kehaiQuoteCodes: createKehaiMap(),
			expect:          map[string]configs.QuoteCodes{},
		},
		{
			name:            "Hassin and Kubun are empty - Return empty map",
			kubun:           model.EmptyString,
			hassin:          model.EmptyString,
			tickQuoteCodes:  createTickMap(),
			kehaiQuoteCodes: createKehaiMap(),
			expect:          map[string]configs.QuoteCodes{},
		},
		{
			name:            "Hassin and Kubun are empty, Kehai Quote Codes and Tick Quote Codes definition are nil - Return empty map",
			kubun:           model.EmptyString,
			hassin:          model.EmptyString,
			tickQuoteCodes:  nil,
			kehaiQuoteCodes: nil,
			expect:          map[string]configs.QuoteCodes{},
		},
		{
			name:            "Tick Quote Codes definition are nil - Return map kehai with key E/T/2",
			kubun:           "E",
			hassin:          "T",
			tickQuoteCodes:  nil,
			kehaiQuoteCodes: createKehaiMap(),
			expect: map[string]configs.QuoteCodes{
				"E/T/2": {
					QKbn:     "E",
					Sndc:     "T",
					LogicID:  "TE10",
					Endpoint: "127.0.0.1",
					DBName:   "tick",
				},
			},
		},
		{
			name:            "Kehai Quote Codes definition are nil - Return map tick with key E/T/1",
			kubun:           "E",
			hassin:          "T",
			tickQuoteCodes:  createTickMap(),
			kehaiQuoteCodes: nil,
			expect: map[string]configs.QuoteCodes{
				"E/T/1": {
					QKbn:     "E",
					Sndc:     "T",
					LogicID:  "TE10",
					Endpoint: "127.0.0.1",
					DBName:   "tick",
				},
			},
		},
		{
			name:            "Kehai Quote Codes and Tick Quote Codes definition are nil - Return empty map",
			kubun:           "E",
			hassin:          "T",
			tickQuoteCodes:  nil,
			kehaiQuoteCodes: nil,
			expect:          map[string]configs.QuoteCodes{},
		},
		{
			name:            "Kehai Quote Codes and Tick Quote Codes definition are nil - Return map with key E/T/1 and E/T/2",
			kubun:           "E",
			hassin:          "T",
			tickQuoteCodes:  createTickMap(),
			kehaiQuoteCodes: createKehaiMap(),
			expect: map[string]configs.QuoteCodes{
				"E/T/1": {
					QKbn:     "E",
					Sndc:     "T",
					LogicID:  "TE10",
					Endpoint: "127.0.0.1",
					DBName:   "tick",
				},
				"E/T/2": {
					QKbn:     "E",
					Sndc:     "T",
					LogicID:  "TE10",
					Endpoint: "127.0.0.1",
					DBName:   "tick",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := service.GetQuoteCodes(test.kubun, test.hassin, test.tickQuoteCodes, test.kehaiQuoteCodes)
			assert.Equal(t, len(test.expect), len(result))
			for key, value := range result {
				assert.Equal(t, test.expect[key].QKbn, value.QKbn)
				assert.Equal(t, test.expect[key].Sndc, value.Sndc)
				assert.Equal(t, test.expect[key].LogicID, value.LogicID)
				assert.Equal(t, test.expect[key].Endpoint, value.Endpoint)
				assert.Equal(t, test.expect[key].DBName, value.DBName)
			}

			assert.Equal(t, result, test.expect)
		})
	}
}

func createTickMap() map[string][]configs.QuoteCodes {
	return map[string][]configs.QuoteCodes{
		"TE1": {
			{
				QKbn:     "E",
				Sndc:     "T",
				LogicID:  "TE10",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
		"TE0": {
			{
				QKbn:     "E",
				Sndc:     "M",
				LogicID:  "TE00",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
			{
				QKbn:     "E",
				Sndc:     "S",
				LogicID:  "TE00",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
		"TF1": {
			{
				QKbn:     "F",
				Sndc:     "T",
				LogicID:  "TE10",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
	}
}

func createKehaiMap() map[string][]configs.QuoteCodes {
	return map[string][]configs.QuoteCodes{
		"CX2": {
			{
				QKbn:     "E",
				Sndc:     "T",
				LogicID:  "TE10",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
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
