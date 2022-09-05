/*
Package loadconfig implements logics load Config from file.
*/
package loadconfig

import "context"

// IConfigurationLoader interface load Config Service
type IConfigurationLoader interface {
	LoadConfig(context.Context) error
}
