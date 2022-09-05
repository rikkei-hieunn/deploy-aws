/*
Package configs contains configuration info.
*/
package configs

import (
	"github.com/spf13/viper"
)

// Server application settings
type Server struct {
	TickSystem
}

// Init application configuration
func Init(path string, fileName string) (*Server, error) {
	cfg := new(Server)

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
