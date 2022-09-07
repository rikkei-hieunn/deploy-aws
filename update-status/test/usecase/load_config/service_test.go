package load_config

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
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
			name: "Load Databases status success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
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
			name: "Load Databases status file not found",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition1.json",
					DevelopEnvironment:                  true,
				},
			}, nil),
			expect: expectLoadDataBaseStatus{
				err:    &fs.PathError{},
				result: configs.GroupDatabaseStatusDefinition{},
			},
		},
		{
			name: "Load Databases status wrong data format",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/wrong_format_file.json",
					DevelopEnvironment:                  true,
				},
			}, nil),
			expect: expectLoadDataBaseStatus{
				err:    &json.UnmarshalTypeError{},
				result: configs.GroupDatabaseStatusDefinition{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.args.LoadDatabaseStatus()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", test.expect.err))
			} else {
				assert.Equal(t, len(test.expect.result.Tick), len(result.Tick))
				for index := range test.expect.result.Tick {
					isExists := false
					for indexActual := range result.Tick {
						if test.expect.result.Tick[index].QKbn == result.Tick[indexActual].QKbn &&
							test.expect.result.Tick[index].Sndc == result.Tick[indexActual].Sndc &&
							test.expect.result.Tick[index].TheSecondKeiStatus == result.Tick[indexActual].TheSecondKeiStatus &&
							test.expect.result.Tick[index].TheFirstKeiStatus == result.Tick[indexActual].TheFirstKeiStatus {
							isExists = true
						}
					}

					if !isExists {
						assert.Fail(t, "Database status does not contain expect list")
					}
				}

				assert.Equal(t, len(test.expect.result.Kehai), len(result.Kehai))
				for index := range test.expect.result.Kehai {
					isExists := false
					for indexActual := range result.Kehai {
						if test.expect.result.Kehai[index].QKbn == result.Kehai[indexActual].QKbn &&
							test.expect.result.Kehai[index].Sndc == result.Kehai[indexActual].Sndc &&
							test.expect.result.Kehai[index].TheSecondKeiStatus == result.Kehai[indexActual].TheSecondKeiStatus &&
							test.expect.result.Kehai[index].TheFirstKeiStatus == result.Kehai[indexActual].TheFirstKeiStatus {
							isExists = true
						}
					}

					if !isExists {
						assert.Fail(t, "Database status does not contain expect list")
					}
				}
			}
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
			name: "Load quote code kei 1 data tick success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
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
			name: "Load quote code kei 1 data kehai success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
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
			name: "Load quote code kei 2 data tick success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
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
			name: "Load quote code kei 2 data kehai success",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei1.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
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
			name: "Load quote code file not found",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/quote_code_definition_tick_kei11.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err:    &fs.PathError{},
				result: nil,
			},
		},
		{
			name: "Load quote code wrong data format",
			args: loadconfig.NewService(&configs.Server{
				TickSystem: configs.TickSystem{
					QuoteCodesDefinitionTickKei1Object:  "definition_files/wrong_format_file.json",
					QuoteCodesDefinitionKehaiKei1Object: "definition_files/quote_code_definition_kehai_kei1.json",
					QuoteCodesDefinitionTickKei2Object:  "definition_files/quote_code_definition_tick_kei2.json",
					QuoteCodesDefinitionKehaiKei2Object: "definition_files/quote_code_definition_kehai_kei2.json",
					DatabaseStatusDefinitionObject:      "definition_files/database_status_definition.json",
					DevelopEnvironment:                  true,
					Kei:                                 "1",
					DataType:                            "1",
				},
			}, nil),
			expect: expectLoadQuoteCodeData{
				err:    &json.UnmarshalTypeError{},
				result: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.args.LoadQuoteCodeData()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", test.expect.err))
			} else {
				assert.Equal(t, len(test.expect.result), len(result))
				for group, quoteCodes := range test.expect.result {
					quoteCodesActual, isExists := result[group]
					if !isExists {
						assert.Fail(t, "Logic group not found")
						return
					}

					for index := range quoteCodes {
						for indexActual := range quoteCodesActual {
							if quoteCodes[index].QKbn == quoteCodesActual[indexActual].QKbn &&
								quoteCodes[index].Sndc == quoteCodesActual[indexActual].Sndc &&
								quoteCodes[index].DBName == quoteCodesActual[indexActual].DBName &&
								quoteCodes[index].Endpoint == quoteCodesActual[indexActual].Endpoint &&
								quoteCodes[index].LogicID == quoteCodesActual[indexActual].LogicID {
								isExists = true
							}
						}

						if !isExists {
							assert.Fail(t, "Logic group not found")
						}
					}

				}
			}
		})
	}
}
