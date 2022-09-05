/*
Package loadconfig all logic for services load config
*/
package loadconfig

import "context"

//ILoadConfig provides all services about load config
type ILoadConfig interface {
	LoadConfig(ctx context.Context,serviceName string) error
}