package configs

import "github.com/spf13/viper"

//Server manege configuration
type Server struct {
	*ECS
	TickSystem
	EnvVarKeys []string
	EnvVarValues []string
}

//Init load config from specific file
func Init(path, filename string) (*Server, error) {
	cfg := new(Server)

	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
