/*
Package configs contains configuration.
*/
package configs

import (
	"github.com/spf13/viper"
	"insert-calendar-infos/model"
)

// Server application settings
type Server struct {
	TickDB
	TickFileBus
	TickSystem
}

// Init application configuration
func Init(path string, fileName string) (*Server, error) {
	cfg := new(Server)

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(model.ContentType)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
