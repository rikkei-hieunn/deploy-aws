package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"start-ecs/api"
	"start-ecs/configs"
	"start-ecs/infrastructure"
	"start-ecs/model"
	"time"
)

const location = "Asia/Tokyo"

var config *configs.Server

// init init configuration
func init() {
	initLoc()
	initLog()
}

var (
	err error
)

func main() {
	ctx := context.Background()

	if err = initConfig(); err != nil {
		log.Error().Msgf("load init config error : %s ", err.Error())

		return
	}

	err = config.Validate()
	if err != nil {
		log.Error().Msgf("validate config error : %s ", err.Error())

		return
	}

	infra, err := infrastructure.Init(ctx, config)
	if err != nil {
		log.Error().Msgf("Error while starting: %s", err.Error())

		return
	}

	server := api.New(infra, config)

	if err = server.Start(ctx); err != nil {
		log.Error().Msgf("start server error : %s ", err.Error())

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

// initConfig get configuration file from S3 and parse to object
func initConfig() error {
	var err error
	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	config, err = configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return err
	}
	config.S3Bucket = bucket
	config.S3Region = region

	return nil
}
