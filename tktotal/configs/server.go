/*
Package configs define all configuration
*/
package configs

import (
	"github.com/spf13/viper"
)

//Server server manage config
type Server struct {
	TickSystem
}

// Init application configuration
func Init(path string, fileName string) (*Server, error) {
	var err error
	cfg := new(Server)

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
