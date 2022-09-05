// Package infrastructure define infra
package infrastructure

import (
	"start-jushin/configs"
	"start-jushin/infrastructure/aws/s3"
	ssm2 "start-jushin/infrastructure/aws/ssm"
)

//Infra instances contain all internal services
type Infra struct {
	SSMHandler ssm2.ISSMHandler
	S3Handler  s3.IS3Handler
}

//Init start connect internal services
func Init(cfg *configs.Server) (*Infra, error) {
	ssmHandler, err := ssm2.NewSSMClient(cfg)
	if err != nil {
		return nil, err
	}
	s3Handler, err := s3.NewS3Storage(&cfg.TickSystem)
	if err != nil {
		return nil, err
	}

	return &Infra{SSMHandler: ssmHandler,
		S3Handler: s3Handler}, nil
}
