package loadconfig

import "context"

// IConfigurationLoader interface load config service
type IConfigurationLoader interface {
	LoadConfig(ctx context.Context) error
}
