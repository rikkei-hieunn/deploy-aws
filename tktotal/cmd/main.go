package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
	"tktotal/api"
	"tktotal/configs"
	"tktotal/infrastructure"
	"tktotal/model"
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

var date string

func main() {
	err := initConfig()

	if err != nil {
		log.Error().Msgf("init config fail : %s", err.Error())

		return
	}

	ctx := context.Background()

	infra, err := infrastructure.Init(ctx, config)
	if err != nil {
		log.Error().Msgf("init infra fail : %s", err.Error())

		return
	}

	server := api.New(infra, config)

	err = server.Start(ctx, date)

	if err != nil {
		log.Error().Msgf("start server fail : %s", err.Error())

		return
	}

	log.Info().Msgf("program executed successfully")
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
	if len(os.Args) > 2 {
		return fmt.Errorf("invalid number params")
	}
	if len(os.Args) == 2 {
		date = os.Args[1]
	}
	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	currentFolder, err := os.Getwd()
	if err != nil {
		return err
	}
	path := currentFolder + model.StrokeCharacter + "environment_variables"
	config, err = configs.Init(path, "environment_variables.json")

	if err != nil {
		return fmt.Errorf("init config fail %w ", err)
	}
	config.S3Bucket = bucket
	config.S3Region = region
	err = config.Validate()
	if err != nil {
		return err
	}

	return nil
}
