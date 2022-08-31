/*
Package infrastructure all logic for internal service
*/
package infrastructure

import (
	"context"
	"start-ecs/configs"
	"start-ecs/infrastructure/aws/ecs"
)

//Infra construct all internal services
type Infra struct {
	ECSHandler ecs.IECSHandler
}

//Init start create infra instance and connect internal services
func Init(ctx context.Context, cfg *configs.Server) *Infra {
	ecsHandler := ecs.NewECSClient(ctx, cfg)

	return &Infra{ECSHandler: ecsHandler}
}