package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"process-get-data/api"
	"process-get-data/configs"
	"process-get-data/infrastructure"
	"process-get-data/model"
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

	processGetDataType := os.Args[1]
	if strings.TrimSpace(processGetDataType) == model.EmptyString {
		return fmt.Errorf("missing arguments")
	}

	if processGetDataType != model.GetDataTypeKei1 && processGetDataType != model.GetDataTypeKei2 && processGetDataType != model.GetDataTypeBoth {
		return fmt.Errorf("invalid process get data type")
	}

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
	config.TickSystem.S3Bucket = bucket
	config.TickSystem.S3Region = region
	config.TickSystem.ProcessGetData = processGetDataType

	return nil
}
