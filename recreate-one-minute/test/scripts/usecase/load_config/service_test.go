package load_config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"recreate-one-minute/configs"
	"recreate-one-minute/model"
	loadconfig "recreate-one-minute/usecase/load_config"
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
	candleManagementPrefix := "candle_management"

	var tests = []struct {
		name   string
		path   string
		expect struct {
			TickDB          *configs.TickDB
			Kubun           map[string]string
			oneMinutePrefix *string
			err             error
		}
	}{
		{
			name: "common definition file not found",
			path: ".",
			expect: struct {
				TickDB          *configs.TickDB
				Kubun           map[string]string
				oneMinutePrefix *string
				err             error
			}{TickDB: nil, Kubun: nil, oneMinutePrefix: nil, err: new(fs.PathError)},
		}, {
			name: "common definition wrong file format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				TickDB          *configs.TickDB
				Kubun           map[string]string
				oneMinutePrefix *string
				err             error
			}{TickDB: nil, Kubun: nil, oneMinutePrefix: nil, err: new(json.UnmarshalTypeError)},
		}, {
			name: "common definition validate error",
			path: "definition_files/common_variables_error.json",
			expect: struct {
				TickDB          *configs.TickDB
				Kubun           map[string]string
				oneMinutePrefix *string
				err             error
			}{TickDB: nil, Kubun: nil, oneMinutePrefix: nil, err: errors.New("database TK_DB_PORT required")},
		}, {
			name: "load common definition file successfully",
			path: "definition_files/common_variables.json",
			expect: struct {
				TickDB          *configs.TickDB
				Kubun           map[string]string
				oneMinutePrefix *string
				err             error
			}{TickDB: &configs.TickDB{
				Port:              3306,
				User:              "admin",
				Password:          "123456123",
				MaxOpenConnection: 100,
				MaxIdleConnection: 100,
				DriverName:        "mysql",
				RetryTimes:        3,
				RetryWaitTimes:    3000,
				Endpoints:         nil,
			}, Kubun: map[string]string{
				"@":  "A1",
				"@@": "A2",
			}, oneMinutePrefix: &candleManagementPrefix, err: nil},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service.Config.CommonDefinitionObject = test.path
			tickDB, candleManagementPrefix, kubunInsteadOf, err := service.LoadCommonData()
			if err != nil {
				assert.Equal(t, test.expect.TickDB, tickDB)
				assert.Equal(t, test.expect.Kubun, kubunInsteadOf)
				assert.Equal(t, test.expect.oneMinutePrefix, candleManagementPrefix)
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, test.expect.TickDB.Port, tickDB.Port)
				assert.Equal(t, test.expect.TickDB.User, tickDB.User)
				assert.Equal(t, test.expect.TickDB.Password, tickDB.Password)
				assert.Equal(t, test.expect.TickDB.MaxOpenConnection, tickDB.MaxOpenConnection)
				assert.Equal(t, test.expect.TickDB.MaxIdleConnection, tickDB.MaxIdleConnection)
				assert.Equal(t, test.expect.TickDB.DriverName, tickDB.DriverName)
				assert.Equal(t, test.expect.TickDB.RetryTimes, tickDB.RetryTimes)
				assert.Equal(t, test.expect.TickDB.RetryWaitTimes, tickDB.RetryWaitTimes)
				assert.Equal(t, len(*test.expect.oneMinutePrefix), len(*candleManagementPrefix))
				assert.Equal(t, len(test.expect.Kubun), len(kubunInsteadOf))
				for key := range kubunInsteadOf {
					assert.Equal(t, test.expect.Kubun[key], kubunInsteadOf[key])
				}
				assert.Equal(t, test.expect.oneMinutePrefix, candleManagementPrefix)
			}
		})
	}
}

func Test_LoadEndpointData(t *testing.T) {
	service := initService()
	var tests = []struct {
		name   string
		path   string
		expect struct {
			endpointMap map[string][]string
			err         error
		}
	}{
		{
			name: "endpoint definition file not found",
			path: ".",
			expect: struct {
				endpointMap map[string][]string
				err         error
			}{endpointMap: nil, err: new(fs.PathError)},
		}, {
			name: "parse endpoint definition wrong format",
			path: "definition_files/wrong_format_file.json",
			expect: struct {
				endpointMap map[string][]string
				err         error
			}{endpointMap: nil, err: new(json.UnmarshalTypeError)},
		}, {
			name: "load endpoint definition file successfully",
			path: "definition_files/quote_code_definition_tick_kei1.json",
			expect: struct {
				endpointMap map[string][]string
				err         error
			}{endpointMap: map[string][]string{
				"localhost/tick": {
					"@/LN", "@/TL", "E/CXJ",
				},
			}, err: nil},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//test on kei1
			service.Config.Kei = model.TheFirstKei
			service.Config.TickSystem.DB1EndpointDefinitionObject = test.path

			endPointMapResult, err := service.LoadEndpointData()
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", test.expect.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, len(test.expect.endpointMap), len(endPointMapResult))
				for endPointKey, quoteCodePairActual := range endPointMapResult {
					endPointKeyExcept := test.expect.endpointMap[endPointKey]
					for i := range endPointMapResult[endPointKey] {
						assert.Equal(t, quoteCodePairActual[i], endPointKeyExcept[i])
					}
				}
			}
		})
	}
}
