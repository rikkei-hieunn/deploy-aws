package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"show-status/api"
	"show-status/configs"
	"show-status/infrastructure"
	"show-status/model"
	"time"
)

const location = "Asia/Tokyo"

// init init configuration
func init() {
	initLoc()
	initLog()
}

func main() {
	config, err := initConfig()
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
func initConfig() (*configs.Server, error) {
	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return nil, fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return nil, fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	config, err := configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return nil, fmt.Errorf("file not found environment_variables/environment_variables.json")
	}
	config.TickSystem.S3Region = region
	config.TickSystem.S3Bucket = bucket

	return config, nil
}
