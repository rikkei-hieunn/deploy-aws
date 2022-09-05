package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"recreate-one-minute/api"
	"recreate-one-minute/configs"
	"recreate-one-minute/infrastructure"
	"recreate-one-minute/model"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	initLoc()
	initLog()
}

var (
	config *configs.Server
)

func main() {
	log.Info().Msg("Start process .....")
	ctx := context.Background()
	err := initConfig()
	if err != nil {
		log.Error().Msgf(" load config fail %s ", err.Error())

		return
	}
	infra, err := infrastructure.Init(config)
	if err != nil {
		log.Error().Msgf("init infra fail %s ", err.Error())

		return
	}
	//create new server instance
	server := api.New(infra, config)
	err = server.Start(ctx)
	if err != nil {
		log.Error().Msgf("restarted instance fail %s ", err.Error())

		return
	}
}
func initLoc() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

// initConfig get configuration file from S3 and parse to object
func initConfig() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("not enough arguments")
	}
	if os.Args[1] != model.TheFirstKei && os.Args[1] != model.TheSecondKei {
		return fmt.Errorf("wrong kei type")
	}
	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	var err error
	config, err = configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return err
	}
	config.TickSystem.TickBucket = bucket
	config.TickSystem.TickRegion = region
	config.TickSystem.Kei = os.Args[1]

	return nil
}
