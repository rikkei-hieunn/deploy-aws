package loadconfig

import "context"

// IConfigurationLoader interface load Config Service
type IConfigurationLoader interface {
	LoadConfig(ctx context.Context) error
}

// IWorker defines worker interfaces
type IWorker interface {
	Start(ctx context.Context)
}
