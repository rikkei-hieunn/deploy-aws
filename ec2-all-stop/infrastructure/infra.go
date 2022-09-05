/*
Package infrastructure implements logics init.
*/
package infrastructure

import (
	"context"
	"ec2-all-stop/configs"
	"ec2-all-stop/infrastructure/aws/ec2"
	"ec2-all-stop/infrastructure/aws/s3"
)

// Infra infrastructure management struct
type Infra struct {
	S3Handler  s3.IS3Handler
	EC2Handler ec2.IEC2Handler
}

// Init initializes resources
func Init(ctx context.Context, cfg *configs.Server) (*Infra, error) {
	ec2Handler, err := ec2.NewEC2Handler(ctx, cfg)
	if err != nil {
		return nil, err
	}

	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{
		EC2Handler: ec2Handler,
		S3Handler:  s3Handler,
	}, nil
}
