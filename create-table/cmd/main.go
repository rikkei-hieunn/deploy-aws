package main

import (
	"context"
	"create-table/api"
	"create-table/configs"
	"create-table/infrastructure"
	"create-table/model"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
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

// initLoc init location, set default is Tokyo
func initLoc() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

// initLog init zero log
func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

// initConfig get configuration file from S3 and parse to object
func initConfig() error {
	var err error
	if len(os.Args) != 2 {
		return fmt.Errorf("missing arguments")
	}

	createType := os.Args[1]
	if strings.TrimSpace(createType) == model.EmptyString {
		return fmt.Errorf("missing arguments")
	}

	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	if createType != model.CreateTypeKei1 && createType != model.CreateTypeKei2 && createType != model.CreateTypeBoth {
		return fmt.Errorf("invalid create type")
	}

	config, err = configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return err
	}
	config.S3Bucket = bucket
	config.S3Region = region
	config.TickSystem.CreateType = createType

	return nil
}
