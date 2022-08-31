/*
Package loadconfig implements logics load config from file.
*/
package loadconfig

import "context"

// IConfigurationLoader interface load config service
type IConfigurationLoader interface {
	LoadConfig(context.Context) error
}
