package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
	"update-status/api"
	"update-status/configs"
	"update-status/infrastructure"
	"update-status/model"
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
	if len(os.Args) < 4 {
		return nil, fmt.Errorf("not enough arguments")
	}

	kei := os.Args[2]
	if kei != model.TheFirstKei && kei != model.TheSecondKei {
		return nil, fmt.Errorf("wrong kei type")
	}

	dataType := os.Args[3]
	if dataType != model.TickData && dataType != model.KehaiData {
		return nil, fmt.Errorf("invalid data type")
	}

	bucket := os.Getenv(model.S3BucketKey)
	if bucket == model.EmptyString {
		return nil, fmt.Errorf("system TK_SYSTEM_BUCKET_NAME required")
	}

	region := os.Getenv(model.S3RegionKey)
	if region == model.EmptyString {
		return nil, fmt.Errorf("system TK_SYSTEM_REGION required")
	}

	var request interface{}
	requestType := os.Args[1]
	// validate number param follow request type
	switch requestType {
	case model.UpdateStatusTypeQuoteCode:
		if len(os.Args) != 7 {
			return nil, fmt.Errorf("not enough arguments")
		}
		boolValue, err := strconv.ParseBool(os.Args[6])
		if err != nil {
			return nil, fmt.Errorf("invalid new status")
		}

		request = model.UpdateTypeQuoteCode{
			Kubun:     os.Args[4],
			Hassin:    os.Args[5],
			NewStatus: boolValue,
		}
	case model.UpdateStatusTypeDBName:
		if len(os.Args) != 6 {
			return nil, fmt.Errorf("not enough arguments")
		}
		boolValue, err := strconv.ParseBool(os.Args[5])
		if err != nil {
			return nil, fmt.Errorf("invalid new status")
		}

		request = model.UpdateTypeDBName{
			DBName:    os.Args[4],
			NewStatus: boolValue,
		}
	case model.UpdateStatusTypeGroupID:
		if len(os.Args) != 6 {
			return nil, fmt.Errorf("not enough arguments")
		}
		boolValue, err := strconv.ParseBool(os.Args[5])
		if err != nil {
			return nil, fmt.Errorf("invalid new status")
		}

		request = model.UpdateTypeGroupID{
			GroupID:   os.Args[4],
			NewStatus: boolValue,
		}
	default:
		return nil, fmt.Errorf("invalid request type")
	}

	config, err := configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return nil, fmt.Errorf("file not found environment_variables/environment_variables.json")
	}
	config.TickSystem.Kei = kei
	config.TickSystem.Request = request
	config.TickSystem.S3Region = region
	config.TickSystem.S3Bucket = bucket
	config.TickSystem.DataType = dataType
	config.TickSystem.RequestType = requestType

	return config, nil
}
