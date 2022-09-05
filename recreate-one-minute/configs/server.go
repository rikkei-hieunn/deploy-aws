package configs

import (
	"github.com/spf13/viper"
)

//Server manage configuration
type Server struct {
	TickDB
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
