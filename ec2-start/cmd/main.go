package main

import (
	"context"
	"ec2-start/api"
	"ec2-start/configs"
	"ec2-start/infrastructure"
	"ec2-start/model"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

const location = "Asia/Tokyo"

var (
	config *configs.Server
)

// init init configuration
func init() {
	initLoc()
	initLog()
}

func main() {
	err := initConfig()
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}

	ctx := context.Background()
	infra, err := infrastructure.Init(ctx, config)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}

	server := api.New(infra, config)
	err = server.Start(ctx)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
}

// initLog init zero log
func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

// initLoc init location, set default is Tokyo
func initLoc() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

// initConfig get list machine ID from os environment
func initConfig() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	// Loop through TickSystem machine ID argument List
	var instanceIDs []string
	for _, machineID := range os.Args[1:] {
		if instanceID := os.Getenv(machineID + model.MachineSuffix); instanceID != model.EmptyString {
			instanceIDs = append(instanceIDs, instanceID)
		}
	}

	if len(instanceIDs) == 0 {
		return fmt.Errorf("no instance's id found")
	}

	config = &configs.Server{
		TickSystem: configs.TickSystem{
			Region:      region,
			InstanceIds: instanceIDs,
		},
	}

	return nil
}
