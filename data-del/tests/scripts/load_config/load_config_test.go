package load_config

import (
	"data-del/configs"
	loadconfig "data-del/usecase/load_config"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestParseCommonData(t *testing.T) {
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
	}{
		{
			name: "success",
			path: "common.json",
		}}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.path)
			if err != nil {
				log.Info().Msg(err.Error())
			}
			table, db, kubun, err := sv.ParseCommonData(b)
			if err != nil {
				return
			}
			assert.Equal(t, reflect.DeepEqual(table, tablePrefixMock), true)
			assert.Equal(t, db.Endpoints, tickDBMock.Endpoints)
			assert.Equal(t, db.RetryTimes, tickDBMock.RetryTimes)
			assert.Equal(t, db.RetryWaitTimes, tickDBMock.RetryWaitTimes)
			assert.Equal(t, db.User, tickDBMock.User)
			assert.Equal(t, db.Password, tickDBMock.Password)
			assert.Equal(t, db.DriverName, tickDBMock.DriverName)
			assert.Equal(t, reflect.DeepEqual(kubun, kubunInsteadOfMock), true)
		})
	}
}
func TestLoadExpiredData(t *testing.T) {
	sv := loadconfig.Service{}
	expiresMock := []configs.TickExpire{
		{QKBN: "Q",
			SNDC:   "Q",
			Expire: 0},
		{QKBN: "E",
			SNDC:   "CXJ",
			Expire: 1},
	}
	expiresAllMock := configs.TickExpire{
		QKBN:   "ALL",
		SNDC:   "",
		Expire: 370,
	}
	t.Run("success", func(t *testing.T) {
		b, err := os.ReadFile("exprired_definition.json")
		if err != nil {
			log.Info().Msg(err.Error())
		}
		expires, expiresAll, err := sv.ParseExpireData(b)
		if err != nil {
			return
		}

		assert.Equal(t, reflect.DeepEqual(expires, expiresMock), true)
		assert.Equal(t, reflect.DeepEqual(expiresAll.Expire, expiresAllMock.Expire), true)
		assert.Equal(t, reflect.DeepEqual(expiresAll.SNDC, expiresAllMock.SNDC), true)
		assert.Equal(t, reflect.DeepEqual(expiresAll.QKBN, expiresAllMock.QKBN), true)
	})
}
func TestParseDBEndpoint(t *testing.T) {
	sv := loadconfig.Service{}
	mock := map[string][]string{
		"127.0.0.1/db-name-1":[]string{"E/T","E/CXJ"},
		"127.0.0.1/db-name-2":[]string{"E/CXJ"},
	}
	t.Run("success", func(t *testing.T) {
		b, err := os.ReadFile("qcd_define.json")
		if err != nil {
			log.Info().Msg(err.Error())
		}
		data, err := sv.ParseDBEndpointData(b)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
		for key, value := range data {
            assert.Equal(t, mock[key],value,true)
		}
	})
}