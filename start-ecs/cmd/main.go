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

	infra := infrastructure.Init(ctx, config)

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
	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}
	dir, err := os.Getwd()
	filePath := dir + model.StrokeCharacter + "environment_variables"
	if err != nil {
		return err
	}
	if config, err = configs.Init(filePath, "environment_variables.json"); err != nil {
		return err
	}
	if len(os.Args) < 4 {
		return fmt.Errorf("not enough arguments")
	}
	config.Region = region

	return nil
}
