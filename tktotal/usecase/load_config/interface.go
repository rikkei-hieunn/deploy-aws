package loadconfig

import "context"

// IConfigurationLoader interface load Config service
type IConfigurationLoader interface {
	LoadConfig(ctx context.Context, date string) error
}
