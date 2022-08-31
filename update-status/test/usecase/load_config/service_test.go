package load_config

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"update-status/configs"
	loadconfig "update-status/usecase/load_config"
)

type expectLoadDataBaseStatus struct {
	result configs.GroupDatabaseStatusDefinition
	err    error
}

func Test_LoadDataBaseStatus(t *testing.T) {
	tests := []struct {
		name   string
		args   loadconfig.IConfigurationLoader
		expect expectLoadDataBaseStatus
	}{
		{
			name: "load config success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
				},
			}, nil),
			expect: expectLoadDataBaseStatus{
				err: nil,
				result: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
				},
			},
		},
		{
			name: "wrong path",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
				},
			}, nil),
			expect: expectLoadDataBaseStatus{
				err:    errors.New("The system cannot find the path specified."),
				result: configs.GroupDatabaseStatusDefinition{},
			},
		},
		{
			name: "missing path",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "",
					DevelopEnvironment:                  true,
				},
			}, nil),
			expect: expectLoadDataBaseStatus{
				err:    errors.New("The system cannot find the file specified."),
				result: configs.GroupDatabaseStatusDefinition{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.args.LoadDatabaseStatus()
			if err != nil {
				assert.Contains(t, err.Error(), test.expect.err.Error())

				return
			}
			actual := expectLoadDataBaseStatus{}
			if result != nil {
				actual.result = *result
			}
			actual.err = err

			assert.Equal(t, test.expect, actual)
		})
	}
}

type expectLoadQuoteCodeData struct {
	result map[string][]configs.QuoteCodes
	err    error
}

func Test_LoadQuoteCodeData(t *testing.T) {
	tests := []struct {
		name   string
		args   loadconfig.IConfigurationLoader
		expect expectLoadQuoteCodeData
	}{
		{
			name: "load config kei 1 data tick success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err: nil,
				result: map[string][]configs.QuoteCodes{
					"AIG": {
						{
							QKbn:     "@",
							Sndc:     "LN",
							LogicID:  "AIG0",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
					"CX2": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX20",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
				},
			},
		},
		{
			name: "load config kei 1 data keihai success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "2",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err: nil,
				result: map[string][]configs.QuoteCodes{
					"CX2": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX20",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
					"JF1": {
						{
							QKbn:     "E",
							Sndc:     "JNF",
							LogicID:  "JF10",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
						{
							QKbn:     "E",
							Sndc:     "JNF",
							LogicID:  "JF11",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
				},
			},
		},
		{
			name: "load config kei 2 data tick success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "2",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err: nil,
				result: map[string][]configs.QuoteCodes{
					"AIG": {
						{
							QKbn:     "S",
							Sndc:     "TW",
							LogicID:  "AIG2",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
					"CX2": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX20",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
				},
			},
		},
		{
			name: "load config kei 2 data keihai success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "2",
					DataType:                            "2",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err: nil,
				result: map[string][]configs.QuoteCodes{
					"CX3": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX30",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
					"CX4": {
						{
							QKbn:     "E",
							Sndc:     "CXJ",
							LogicID:  "CX40",
							Endpoint: "127.0.0.1",
							DBName:   "tick",
						},
					},
				},
			},
		},
		{
			name: "load config wrong path",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "../environment_variables/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err:    errors.New("The system cannot find the path specified."),
				result: nil,
			},
		},
		{
			name: "load config missing path",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "",
					QuoteCodesDefinitionKehaiKei1Object: "../../environment_variables/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "../../environment_variables/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "../../environment_variables/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "../../environment_variables/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err:    errors.New("The system cannot find the file specified."),
				result: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.args.LoadQuoteCodeData()

			actual := expectLoadQuoteCodeData{
				err:    err,
				result: result,
			}

			if err != nil {
				assert.Contains(t, err.Error(), test.expect.err.Error())

				return
			}

			assert.Equal(t, test.expect, actual)
		})
	}
}
