package load_config

import (
	"data-del/configs"
	loadconfig "data-del/usecase/load_config"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func Test_ParseCommonData(t *testing.T) {
	sv := loadconfig.Service{}
	tickDBMock := configs.TickDB{
		Port:              3306,
		User:              "admin",
		Password:          "123456123",
		MaxIdleConnection: 100,
		MaxOpenConnection: 100,
		DriverName:        "mysql",
		RetryTimes:        3,
		RetryWaitTimes:    3000,
	}
	tablePrefixMock := map[int]string{
		0: "business_day_information_data",
		1: "best_quote_data",
	}
	kubunInsteadOfMock := map[string]string{
		"@":  "A1",
		"@@": "A2",
	}
	var test = []struct {
		name string
		path string
		args []byte
		err  error
	}{
		{
			name: "Parse common data success",
			path: "definition_files/common.json",
			args: []byte{},
			err:  nil,
		},
		{
			name: "Parse common data wrong data format",
			path: "definition_files/wrong_format_file.json",
			args: []byte{},
			err:  &json.UnmarshalTypeError{},
		},
		{
			name: "Parse common data - input data nil",
			args: nil,
			err:  &json.SyntaxError{},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args != nil {
				b, err := os.ReadFile(tt.path)
				if err != nil {
					log.Info().Msg(err.Error())
					return
				}
				tt.args = append(tt.args, b...)
			}
			table, db, kubun, err := sv.ParseCommonData(tt.args)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, reflect.DeepEqual(table, tablePrefixMock), true)
				assert.Equal(t, db.Endpoints, tickDBMock.Endpoints)
				assert.Equal(t, db.RetryTimes, tickDBMock.RetryTimes)
				assert.Equal(t, db.RetryWaitTimes, tickDBMock.RetryWaitTimes)
				assert.Equal(t, db.User, tickDBMock.User)
				assert.Equal(t, db.Password, tickDBMock.Password)
				assert.Equal(t, db.DriverName, tickDBMock.DriverName)
				assert.Equal(t, reflect.DeepEqual(kubun, kubunInsteadOfMock), true)
			}
		})
	}
}

func Test_LoadExpiredData(t *testing.T) {
	sv := loadconfig.Service{}
	expiresMock := []configs.TickExpire{
		{
			QKBN:   "Q",
			SNDC:   "Q",
			Expire: 0,
		},
		{
			QKBN:   "E",
			SNDC:   "CXJ",
			Expire: 1,
		},
	}
	expiresAllMock := configs.TickExpire{
		QKBN:   "ALL",
		SNDC:   "",
		Expire: 370,
	}
	var test = []struct {
		name string
		path string
		args []byte
		err  error
	}{
		{
			name: "Load expired data success",
			path: "definition_files/exprired_definition.json",
			args: []byte{},
			err:  nil,
		},
		{
			name: "Load expired data wrong data format",
			path: "definition_files/wrong_format_file.json",
			args: []byte{},
			err:  &json.UnmarshalTypeError{},
		},
		{
			name: "Load expired data - input data nil",
			args: nil,
			err:  &json.SyntaxError{},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args != nil {
				b, err := os.ReadFile(tt.path)
				if err != nil {
					log.Info().Msg(err.Error())
					return
				}
				tt.args = append(tt.args, b...)
			}
			expires, expiresAll, err := sv.ParseExpireData(tt.args)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, reflect.DeepEqual(expires, expiresMock), true)
				assert.Equal(t, reflect.DeepEqual(expiresAll.Expire, expiresAllMock.Expire), true)
				assert.Equal(t, reflect.DeepEqual(expiresAll.SNDC, expiresAllMock.SNDC), true)
				assert.Equal(t, reflect.DeepEqual(expiresAll.QKBN, expiresAllMock.QKBN), true)
			}
		})
	}
}

func Test_ParseDBEndpoint(t *testing.T) {
	sv := loadconfig.Service{}
	mock := map[string][]string{
		"127.0.0.1/db-name-1": {"E/T", "E/CXJ"},
		"127.0.0.1/db-name-2": {"E/CXJ"},
	}
	var test = []struct {
		name string
		path string
		args []byte
		err  error
	}{
		{
			name: "Parse DB endpoint success",
			path: "definition_files/qcd_define.json",
			args: []byte{},
			err:  nil,
		},
		{
			name: "Parse DB endpoint wrong data format",
			path: "definition_files/wrong_format_file.json",
			args: []byte{},
			err:  &json.UnmarshalTypeError{},
		},
		{
			name: "Parse DB endpoint - input data nil",
			args: nil,
			err:  &json.SyntaxError{},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args != nil {
				b, err := os.ReadFile(tt.path)
				if err != nil {
					log.Info().Msg(err.Error())
					return
				}
				tt.args = append(tt.args, b...)
			}
			data, err := sv.ParseDBEndpointData(tt.args)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", err), fmt.Sprintf("%T", tt.err))
			} else {
				assert.Equal(t, len(mock), len(data))
				for key, value := range data {
					assert.Equal(t, mock[key], value, true)
				}
			}
		})
	}
}
