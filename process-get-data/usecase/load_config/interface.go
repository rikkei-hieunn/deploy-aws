package loadconfig

// IConfigurationLoader interface load Config Service
type IConfigurationLoader interface {
	LoadConfig() error
}
