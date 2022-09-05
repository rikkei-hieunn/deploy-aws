package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
	"toiawase-start/api"
	"toiawase-start/configs"
	"toiawase-start/infrastructure"
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

// initConfig get configuration file from S3 and parse to object
func initConfig() error {
	var err error

	config, err = configs.Init()
	if err != nil {
		return err
	}

	return nil
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
