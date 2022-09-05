/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"context"
	"ec2-stop/configs"
	"ec2-stop/infrastructure/aws/ec2"
)

// Infra infrastructure management struct
type Infra struct {
	EC2Handler ec2.IEC2Handler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	ec2Handler, err := ec2.NewEC2Handler(ctx, &cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{
		EC2Handler: ec2Handler,
	}, nil
}
