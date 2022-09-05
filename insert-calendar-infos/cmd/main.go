package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"insert-calendar-infos/api"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/infrastructure"
	"insert-calendar-infos/model"
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

	if err = config.TickSystem.Validate(); err != nil {
		log.Error().Msg(err.Error())

		return
	}

	infra, err := infrastructure.Init(config)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}

	server := api.New(infra, config)
	err = server.Start(context.Background())
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
	defer func() {
		server.Close()
	}()
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
