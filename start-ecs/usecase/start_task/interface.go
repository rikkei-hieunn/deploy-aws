/*
Package starttask all logic for services to start task ecs
*/
package starttask

import "context"

//IStartTask provides run task services
type IStartTask interface {
	Start(ctx context.Context) error
}
