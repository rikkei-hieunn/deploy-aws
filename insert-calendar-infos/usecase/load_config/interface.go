/*
Package loadconfig implements logics load config from file.
*/
package loadconfig

// IConfigurationLoader interface load config service
type IConfigurationLoader interface {
	LoadConfig() error
}
