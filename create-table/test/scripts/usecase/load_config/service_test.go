package load_config

import (
	"create-table/configs"
	"create-table/model"
	loadconfig "create-table/usecase/load_config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
	"time"
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
	oneMinutePrefix := "one_minute_data"

	var tests = []struct {
		name   string
		path   string
		expect struct {
			configDB        *configs.TickDB
			tableNamePrefix *configs.TableNamePrefix
			kubuns          map[string]string
			oneMinutePrefix *string
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
				oneMinutePrefix *string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				oneMinutePrefix: nil,
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
				oneMinutePrefix *string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				oneMinutePrefix: nil,
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
				oneMinutePrefix *string
				err             error
			}{
				configDB:        nil,
				tableNamePrefix: nil,
				kubuns:          nil,
				oneMinutePrefix: nil,
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
				oneMinutePrefix *string
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
				oneMinutePrefix: &oneMinutePrefix,
				err:             nil,
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
				assert.Equal(t, test.expect.oneMinutePrefix, oneMinPrefix)
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
				assert.Equal(t, test.expect.oneMinutePrefix, oneMinPrefix)
			}
		})
	}
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
						LogicGroup:  "TE1",
						TopicName:   "SZN-TSE1",
						Types:       "1,2,4",
						TypesString: []string{"1", "2", "4"},
					},
					{
						LogicGroup:  "TE2",
						TopicName:   "SZN-TSE2",
						Types:       "1,2,4",
						TypesString: []string{"1", "2", "4"},
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.GroupsDefinitionObject = test.path
			result, err := service.LoadGroupData()
			if err != nil {
				assert.Equal(t, test.expect.configDB, result)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.err, err)
				assert.Equal(t, len(test.expect.configDB), len(result))
				for i := range result {
					assert.Equal(t, test.expect.configDB[i].TopicName, result[i].TopicName)
					assert.Equal(t, test.expect.configDB[i].LogicGroup, result[i].LogicGroup)
					assert.Equal(t, test.expect.configDB[i].Types, result[i].Types)
					assert.Equal(t, len(test.expect.configDB[i].TypesString), len(result[i].TypesString))
					for j := range result[i].TypesString {
						assert.Equal(t, test.expect.configDB[i].TypesString[j], result[i].TypesString[j])
					}
				}
			}
		})
	}
}

func Test_LoadElementsData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			tableDefinitions map[string]map[string]configs.TableDefinition
			err              error
		}
	}{
		{
			name: "elements definition file not found",
			path: ".",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: nil,
				err:              new(fs.PathError),
			},
		},
		{
			name: "parse elements definition file wrong format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: nil,
				err:              new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse elements definition wrong start date format",
			path: "definition_files/elements_definition_wrong_start_date.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: nil,
				err:              new(time.ParseError),
			},
		},
		{
			name: "parse elements definition wrong end date format",
			path: "definition_files/elements_definition_wrong_end_date.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: nil,
				err:              new(time.ParseError),
			},
		},
		{
			name: "parse elements definition file start date after now",
			path: "definition_files/elements_definition_start_date_after_now.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: map[string]map[string]configs.TableDefinition{},
				err:              nil,
			},
		},
		{
			name: "parse elements definition file successfully - start date before now - end date empty",
			path: "definition_files/elements_definition_start_date_before_now_end_date_empty.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: map[string]map[string]configs.TableDefinition{
					"@": {
						"1": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"-1": {
							Elements: []configs.Element{
								{
									Name:   "TIME",
									Column: "TIME",
									Length: 5,
								},
							},
							StartDate: "2022/05/01",
							EndDate:   "2032/05/01",
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "parse one minute definition file successfully - start date before now - end date before now",
			path: "definition_files/elements_definition_start_date_before_now_end_date_before_now.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: map[string]map[string]configs.TableDefinition{},
				err:              nil,
			},
		},
		{
			name: "parse elements definition file successfully - start date before now - end date after now",
			path: "definition_files/elements_definition_start_date_before_now_end_date_after_now.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: map[string]map[string]configs.TableDefinition{
					"@": {
						"1": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "3022/05/10",
						},
						"-1": {
							Elements: []configs.Element{
								{
									Name:   "TIME",
									Column: "TIME",
									Length: 5,
								},
							},
							StartDate: "2022/05/01",
							EndDate:   "2032/05/01",
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "parse all elements definition file successfully",
			path: "definition_files/elements_definition.json",
			expect: struct {
				tableDefinitions map[string]map[string]configs.TableDefinition
				err              error
			}{
				tableDefinitions: map[string]map[string]configs.TableDefinition{
					"@": {
						"1": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"2": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"3": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"4": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"5": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "",
						},
						"6": {
							Elements: []configs.Element{
								{
									Name:   "QCD",
									Column: "QCD",
									Length: 42,
								},
							},
							StartDate: "2022/05/10",
							EndDate:   "3022/05/10",
						},
						"-1": {
							Elements: []configs.Element{
								{
									Name:   "TIME",
									Column: "TIME",
									Length: 5,
								},
							},
							StartDate: "2022/05/01",
							EndDate:   "2032/05/01",
						},
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.ElementsDefinitionObject = test.path
			result, err := service.LoadElementsData(createOneMinuteTable())
			if err != nil {
				assert.Equal(t, test.expect.tableDefinitions, result)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.err, err)
				assert.Equal(t, len(test.expect.tableDefinitions), len(result))
				for i := range result {
					tableDefinitionMapExpect := test.expect.tableDefinitions[i]
					tableDefinitionMapActual := result[i]
					assert.Equal(t, len(tableDefinitionMapExpect), len(tableDefinitionMapActual))
					for j := range tableDefinitionMapActual {
						assert.Equal(t, len(tableDefinitionMapExpect[j].Elements), len(tableDefinitionMapActual[j].Elements))
						for k := range tableDefinitionMapActual[j].Elements {
							assert.Equal(t, tableDefinitionMapExpect[j].Elements[k].Name, tableDefinitionMapActual[j].Elements[k].Name)
							assert.Equal(t, tableDefinitionMapExpect[j].Elements[k].Column, tableDefinitionMapActual[j].Elements[k].Column)
							assert.Equal(t, tableDefinitionMapExpect[j].Elements[k].Length, tableDefinitionMapActual[j].Elements[k].Length)
						}
						assert.Equal(t, tableDefinitionMapExpect[j].StartDate, tableDefinitionMapActual[j].StartDate)
						assert.Equal(t, tableDefinitionMapExpect[j].EndDate, tableDefinitionMapActual[j].EndDate)
					}
				}
			}
		})
	}
}

func Test_LoadOneMinuteElementsData(t *testing.T) {
	service := initService()

	var tests = []struct {
		name   string
		path   string
		expect struct {
			oneMinTables map[string]configs.OneMinuteTableDefinition
			err          error
		}
	}{
		{
			name: "one minute definition file not found",
			path: ".",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: nil,
				err:          new(fs.PathError),
			},
		},
		{
			name: "parse one minute definition file wrong format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: nil,
				err:          new(json.UnmarshalTypeError),
			},
		},
		{
			name: "parse one minute definition wrong start date format",
			path: "definition_files/one_minute_columns_definition_wrong_start_date.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: nil,
				err:          new(time.ParseError),
			},
		},
		{
			name: "parse one minute definition wrong end date format",
			path: "definition_files/one_minute_columns_definition_wrong_end_date.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: nil,
				err:          new(time.ParseError),
			},
		},
		{
			name: "parse one minute definition file start date after now",
			path: "definition_files/one_minute_columns_definition_start_date_after_now.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: map[string]configs.OneMinuteTableDefinition{},
				err:          nil,
			},
		},
		{
			name: "parse one minute definition file successfully - start date before now - end date empty",
			path: "definition_files/one_minute_columns_definition_start_date_before_now_end_date_empty.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: map[string]configs.OneMinuteTableDefinition{
					"E": {
						TkQkbn: "E",
						Elements: []configs.Element{
							{
								Name:   "QCD",
								Column: "QCD",
								Length: 42,
							},
						},
						StartDate: "2020/05/10",
						EndDate:   "",
					},
				},
				err: nil,
			},
		},
		{
			name: "parse one minute definition file successfully - start date before now - end date before now",
			path: "definition_files/one_minute_columns_definition_start_date_before_now_end_date_before_now.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: map[string]configs.OneMinuteTableDefinition{},
				err:          nil,
			},
		},
		{
			name: "parse one minute definition file successfully - start date before now - end date after now",
			path: "definition_files/one_minute_columns_definition_start_date_before_now_end_date_after_now.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: map[string]configs.OneMinuteTableDefinition{
					"E": {
						TkQkbn: "E",
						Elements: []configs.Element{
							{
								Name:   "QCD",
								Column: "QCD",
								Length: 42,
							},
						},
						StartDate: "2020/05/10",
						EndDate:   "3020/05/12",
					},
				},
				err: nil,
			},
		},
		{
			name: "parse one minute definition file successfully",
			path: "definition_files/one_minute_columns_definition.json",
			expect: struct {
				oneMinTables map[string]configs.OneMinuteTableDefinition
				err          error
			}{
				oneMinTables: map[string]configs.OneMinuteTableDefinition{
					"E": {
						TkQkbn: "E",
						Elements: []configs.Element{
							{
								Name:   "QCD",
								Column: "QCD",
								Length: 42,
							},
							{
								Name:   "TIME",
								Column: "TIME",
								Length: 5,
							},
						},
						StartDate: "2022/05/01",
						EndDate:   "2030/05/10",
					},
					"F": {
						TkQkbn: "F",
						Elements: []configs.Element{
							{
								Name:   "DFLAG",
								Column: "DFLAG",
								Length: 1,
							},
						},
						StartDate: "2022/05/01",
						EndDate:   "2030/05/10",
					},
				},
				err: nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.OneMinuteDefinitionObject = test.path
			result, err := service.LoadOneMinuteElementsData()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, len(test.expect.oneMinTables), len(result))
				for key, oneMinActual := range result {
					oneMinExpect := test.expect.oneMinTables[key]
					assert.NotEqual(t, oneMinActual, nil)
					assert.Equal(t, oneMinExpect.TkQkbn, oneMinActual.TkQkbn)
					assert.Equal(t, oneMinExpect.StartDate, oneMinActual.StartDate)
					assert.Equal(t, oneMinExpect.EndDate, oneMinActual.EndDate)
					assert.Equal(t, len(oneMinExpect.Elements), len(oneMinActual.Elements))
					for i, _ := range oneMinActual.Elements {
						assert.Equal(t, oneMinExpect.Elements[i].Name, oneMinActual.Elements[i].Name)
						assert.Equal(t, oneMinExpect.Elements[i].Column, oneMinActual.Elements[i].Column)
						assert.Equal(t, oneMinExpect.Elements[i].Length, oneMinActual.Elements[i].Length)
					}
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

func Test_ParseTargetCreateTable(t *testing.T) {
	service := new(loadconfig.Service)

	model.TablePrefix = map[string]string{
		"1": "best_quote_data",
		"2": "multiple_quote_data",
		"3": "master_quote_data",
		"4": "money_flow_data",
		"5": "option_extended_data",
		"6": "index_trading_data",
	}

	model.OneMinuteTablePrefix = "one_minute_data"

	var tests = []struct {
		name            string
		groups          []configs.Group
		oneMinuteTables map[string]configs.OneMinuteTableDefinition
		expect          []configs.TargetCreateTable
	}{
		{
			"No group matched",
			[]configs.Group{
				{
					LogicGroup:  "TE6",
					TopicName:   "SZN-TSE6",
					Types:       "1,2,4",
					TypesString: []string{"1", "2", "4"},
				},
				{
					LogicGroup:  "TE7",
					TopicName:   "SZN-TSE7",
					Types:       "1,2,4",
					TypesString: []string{"1", "2", "4"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{},
		},
		{
			"Groups matched exist - best_quote_data",
			[]configs.Group{
				{
					LogicGroup:  "TE1",
					TopicName:   "SZN-TSE1",
					Types:       "1",
					TypesString: []string{"1"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TE1",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "1",
					TablePrefix: "best_quote_data",
				},
			},
		},
		{
			"Groups matched exist - multiple_quote_data",
			[]configs.Group{
				{
					LogicGroup:  "TE1",
					TopicName:   "SZN-TSE1",
					Types:       "2",
					TypesString: []string{"2"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TE1",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "2",
					TablePrefix: "multiple_quote_data",
				},
			},
		},
		{
			"Groups matched exist - master_quote_data",
			[]configs.Group{
				{
					LogicGroup:  "CX2",
					TopicName:   "SZN-TSE1",
					Types:       "3",
					TypesString: []string{"3"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "CX2",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "3",
					TablePrefix: "master_quote_data",
				},
			},
		},
		{
			"Groups matched exist - money_flow_data",
			[]configs.Group{
				{
					LogicGroup:  "TE1",
					TopicName:   "SZN-TSE1",
					Types:       "4",
					TypesString: []string{"4"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TE1",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "4",
					TablePrefix: "money_flow_data",
				},
			},
		},
		{
			"Groups matched exist - option_extended_data",
			[]configs.Group{
				{
					LogicGroup:  "TE1",
					TopicName:   "SZN-TSE1",
					Types:       "5",
					TypesString: []string{"5"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TE1",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "5",
					TablePrefix: "option_extended_data",
				},
			},
		},
		{
			"Groups matched exist - index_trading_data",
			[]configs.Group{
				{
					LogicGroup:  "TE1",
					TopicName:   "SZN-TSE1",
					Types:       "6",
					TypesString: []string{"6"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TE1",
					QKbn:        "E",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "6",
					TablePrefix: "index_trading_data",
				},
			},
		},
		{
			"Groups matched exist - one_minute_data",
			[]configs.Group{
				{
					LogicGroup:  "TF1",
					TopicName:   "SZN-TSE1",
					Types:       "-1",
					TypesString: []string{"-1"},
				},
			},
			createOneMinuteTable(),
			[]configs.TargetCreateTable{
				{
					LogicGroup:  "TF1",
					QKbn:        "F",
					Sndc:        "T",
					Endpoint:    "127.0.0.1",
					DBName:      "tick",
					DataType:    "-1",
					TablePrefix: "one_minute_data",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := service.ParseTargetCreateTable(createTickMap(), createKehaiMap(), test.groups, createOneMinuteTable())
			assert.Equal(t, len(result), len(test.expect))
			if len(result) != 0 {
				for i := range result {
					assert.Equal(t, test.expect[i].QKbn, result[i].QKbn)
					assert.Equal(t, test.expect[i].Sndc, result[i].Sndc)
					assert.Equal(t, test.expect[i].LogicGroup, result[i].LogicGroup)
					assert.Equal(t, test.expect[i].DataType, result[i].DataType)
					assert.Equal(t, result[i].TablePrefix, test.expect[i].TablePrefix)
					assert.Equal(t, result[i].Endpoint, test.expect[i].Endpoint)
					assert.Equal(t, result[i].DBName, test.expect[i].DBName)
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

func createOneMinuteTable() map[string]configs.OneMinuteTableDefinition {
	return map[string]configs.OneMinuteTableDefinition{
		"F": {
			TkQkbn: "F",
			Elements: []configs.Element{
				{
					Name:   "QCD",
					Column: "QCD",
					Length: 42,
				},
				{
					Name:   "TIME",
					Column: "TIME",
					Length: 5,
				},
			},
			StartDate: "2022/05/01",
			EndDate:   "2032/05/01",
		},
		"@": {
			TkQkbn: "@",
			Elements: []configs.Element{
				{
					Name:   "TIME",
					Column: "TIME",
					Length: 5,
				},
			},
			StartDate: "2022/05/01",
			EndDate:   "2032/05/01",
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
