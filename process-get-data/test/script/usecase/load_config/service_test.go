package load_config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"process-get-data/configs"
	loadconfig "process-get-data/usecase/load_config"
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
	candleManagementTableName := "candle_management"

	var tests = []struct {
		name   string
		path   string
		expect struct {
			configDB         *configs.TickDB
			tableNamePrefix  *configs.TableNamePrefix
			kubuns           map[string]string
			candleManagement *string
			err              error
		}
	}{
		{
			name: "common definition file not found",
			path: ".",
			expect: struct {
				configDB         *configs.TickDB
				tableNamePrefix  *configs.TableNamePrefix
				kubuns           map[string]string
				candleManagement *string
				err              error
			}{
				configDB:         nil,
				tableNamePrefix:  nil,
				kubuns:           nil,
				candleManagement: nil,
				err:              new(fs.PathError),
			},
		},
		{
			name: "common definition wrong file format",
			path: "./definition_files/wrong_format_file.json",
			expect: struct {
				configDB         *configs.TickDB
				tableNamePrefix  *configs.TableNamePrefix
				kubuns           map[string]string
				candleManagement *string
				err              error
			}{
				configDB:         nil,
				tableNamePrefix:  nil,
				kubuns:           nil,
				candleManagement: nil,
				err:              new(json.UnmarshalTypeError),
			},
		},
		{
			name: "common definition validate error",
			path: "./definition_files/common_variables_error.json",
			expect: struct {
				configDB         *configs.TickDB
				tableNamePrefix  *configs.TableNamePrefix
				kubuns           map[string]string
				candleManagement *string
				err              error
			}{
				configDB:         nil,
				tableNamePrefix:  nil,
				kubuns:           nil,
				candleManagement: nil,
				err:              errors.New("database TK_DB_PORT required"),
			},
		},
		{
			name: "load common definition file successfully",
			path: "./definition_files/common_variables.json",
			expect: struct {
				configDB         *configs.TickDB
				tableNamePrefix  *configs.TableNamePrefix
				kubuns           map[string]string
				candleManagement *string
				err              error
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
				candleManagement: &candleManagementTableName,
				err:              nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.CommonDefinitionObject = test.path
			tickDB, tablePrefixes, kubunMap, oneMinPrefix, err := service.LoadCommonData()
			if err != nil {
				assert.Equal(t, test.expect.configDB, tickDB)
				assert.Equal(t, test.expect.tableNamePrefix, tablePrefixes)
				assert.Equal(t, test.expect.kubuns, kubunMap)
				assert.Equal(t, test.expect.candleManagement, oneMinPrefix)
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
					assert.Equal(t, (*test.expect.tableNamePrefix)[i].DataType, (*tablePrefixes)[i].DataType)
					assert.Equal(t, (*test.expect.tableNamePrefix)[i].Prefix, (*tablePrefixes)[i].Prefix)
				}
				assert.Equal(t, len(test.expect.kubuns), len(kubunMap))
				for key := range kubunMap {
					assert.Equal(t, test.expect.kubuns[key], kubunMap[key])
				}
				assert.Equal(t, *test.expect.candleManagement, *oneMinPrefix)
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
			quoteCodes map[string]configs.QuoteCodes
			err        error
		}
	}{
		{
			name: "quote code definition file not found",
			path: ".",
			expect: struct {
				quoteCodes map[string]configs.QuoteCodes
				err        error
			}{
				quoteCodes: nil,
				err:        new(fs.PathError),
			},
		},
		{
			name: "parse quote code definition file wrong format",
			path: "./definition_files/wrong_format_file.json",
			expect: struct {
				quoteCodes map[string]configs.QuoteCodes
				err        error
			}{
				quoteCodes: nil,
				err:        new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse quote code definition file successfully",
			path: "./definition_files/quote_code_definition_tick_kei1.json",
			expect: struct {
				quoteCodes map[string]configs.QuoteCodes
				err        error
			}{
				quoteCodes: map[string]configs.QuoteCodes{
					"@/LN": {
						QKbn:     "@",
						Sndc:     "LN",
						LogicID:  "AIG0",
						Endpoint: "127.0.0.1",
						DBName:   "tick",
					},
					"@/TL": {
						QKbn:     "@",
						Sndc:     "TL",
						LogicID:  "AIG0",
						Endpoint: "127.0.0.1",
						DBName:   "tick",
					},
					"E/CXJ": {
						QKbn:     "E",
						Sndc:     "CXJ",
						LogicID:  "CX30",
						Endpoint: "127.0.0.1",
						DBName:   "tick",
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
					assert.NotEqual(t, quoteCodeExpect, nil)
					assert.Equal(t, quoteCodeExpect.QKbn, quoteCodeActual.QKbn)
					assert.Equal(t, quoteCodeExpect.Sndc, quoteCodeActual.Sndc)
					assert.Equal(t, quoteCodeActual.LogicID, quoteCodeExpect.LogicID)
					assert.Equal(t, quoteCodeActual.Endpoint, quoteCodeExpect.Endpoint)
					assert.Equal(t, quoteCodeActual.DBName, quoteCodeExpect.DBName)
				}
			}
		})
	}
}

func Test_LoadOneMinuteConfigData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			oneMinuteConfigs []configs.OneMinuteConfig
			err              error
		}
	}{
		{
			name: "one minute operation configs file not found",
			path: ".",
			expect: struct {
				oneMinuteConfigs []configs.OneMinuteConfig
				err              error
			}{
				oneMinuteConfigs: nil,
				err:              new(fs.PathError),
			},
		},
		{
			name: "parse one minute operation configs file wrong format",
			path: "./definition_files/wrong_format_file.json",
			expect: struct {
				oneMinuteConfigs []configs.OneMinuteConfig
				err              error
			}{
				oneMinuteConfigs: nil,
				err:              new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse one minute operation configs file successfully",
			path: "./definition_files/one_minute_operation_configs.json",
			expect: struct {
				oneMinuteConfigs []configs.OneMinuteConfig
				err              error
			}{
				oneMinuteConfigs: []configs.OneMinuteConfig{
					{
						QKbn: "E",
						Sndc: "T",
						OperatorType: "0",
						OriginStartIndex: "0",
						StartIndex: "0",
						EndIndex: "1",
						QuoteCode: "",
						CreateTime: "12:00",
						CreateDay: 0,
						StartTime: "08:00:00",
						EndTime: "11:30:59",
						TableName: "one_minute_data_E_T",
						Mon: "1",
						Tue: "0",
						Wed: "1",
						Thu: "1",
						Fri: "1",
						Sat: "0",
						Sun: "1",
					},
					{
						QKbn: "E",
						Sndc: "O",
						OperatorType: "0",
						OriginStartIndex: "",
						StartIndex: "-1",
						EndIndex: "",
						QuoteCode: "",
						CreateTime: "12:00",
						CreateDay: -1,
						StartTime: "08:00:00",
						EndTime: "11:30:59",
						TableName: "one_minute_data_E_O",
						Mon: "0",
						Tue: "0",
						Wed: "1",
						Thu: "1",
						Fri: "1",
						Sat: "0",
						Sun: "1",
					},
				},
				err:              nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.OneMinuteOperatorConfigObject = test.path
			result, err := service.LoadOneMinuteConfigData()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, len(test.expect.oneMinuteConfigs), len(result))
				for i, _ := range result {
					assert.Equal(t, test.expect.oneMinuteConfigs[i].QKbn, result[i].QKbn)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Sndc, result[i].Sndc)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].OperatorType, result[i].OperatorType)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].OriginStartIndex, result[i].OriginStartIndex)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].StartIndex, result[i].StartIndex)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].EndIndex, result[i].EndIndex)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].QuoteCode, result[i].QuoteCode)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].CreateTime, result[i].CreateTime)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].CreateDay, result[i].CreateDay)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].StartTime, result[i].StartTime)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].EndTime, result[i].EndTime)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].TableName, result[i].TableName)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Mon, result[i].Mon)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Tue, result[i].Tue)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Wed, result[i].Wed)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Thu, result[i].Thu)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Fri, result[i].Fri)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Sat, result[i].Sat)
					assert.Equal(t, test.expect.oneMinuteConfigs[i].Sun, result[i].Sun)
				}
			}
		})
	}
}
