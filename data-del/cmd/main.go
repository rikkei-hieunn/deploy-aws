package main

import (
	"context"
	"data-del/api"
	"data-del/configs"
	"data-del/infrastructure"
	"data-del/model"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

const location = "Asia/Tokyo"

// init init configuration
func init() {
	initLoc()
	initLog()
}

// initLog init zero log
func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

var (
	err       error
	config    *configs.Server
	keiNumber string
)

// initLoc init location, set default is Tokyo
func initLoc() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	ctx := context.Background()
	//init configuration
	err := initConfig()
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
	//init infra
	infra, err := infrastructure.Init(ctx, config)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
	//create server
	server := api.New(infra, config)

	err = server.Start(ctx, keiNumber)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
	defer func() {
		server.Close()
	}()
}

//init configuration
func initConfig() error {
	var deletedDaysNumber int
	if len(os.Args) == 2 {
		keiNumber = os.Args[1]
		//default is 1
		deletedDaysNumber = 1
	} else if len(os.Args) == 3 {
		keiNumber = os.Args[1]
		deletedDaysNumber, err = strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid arguments days : %s ", os.Args[2])
		}
	} else {
		return fmt.Errorf("missing arguments")
	}
	if keiNumber != model.FirstKei && keiNumber != model.SecondKei {
		return fmt.Errorf("invalid kei %s", keiNumber)
	}
	if deletedDaysNumber < 0 {
		return fmt.Errorf("invalid delete days number %d", deletedDaysNumber)
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

	config.NumberOfDeletedDays = deletedDaysNumber
	config.S3Bucket = bucket
	config.S3Region = region
	err = config.TickSystem.Validate()
	if err != nil {
		log.Error().Msg(err.Error())

		return err
	}

	return nil
}
