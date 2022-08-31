package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"send-command/api"
	"send-command/configs"
	"send-command/infrastructure"
	"send-command/model"
	"time"
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
	if len(os.Args) < 3 {
		return nil, fmt.Errorf("not enough arguments")
	}
	if os.Args[2] != model.TheFirstKei && os.Args[2] != model.TheSecondKei {
		return nil, fmt.Errorf("wrong kei type")
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
	case model.RequestTypeAll:
		if len(os.Args) != 4 {
			return nil, fmt.Errorf("not enough arguments")
		}
		request = model.RequestAll{
			Kei:     os.Args[2],
			Command: os.Args[3],
		}
	case model.RequestTypeGroupLine:
		if len(os.Args) != 5 {
			return nil, fmt.Errorf("not enough arguments")
		}
		request = model.RequestLine{
			Kei:       os.Args[2],
			GroupLine: os.Args[3],
			Command:   os.Args[4],
		}
	case model.RequestTypeGroupID:
		if len(os.Args) != 6 {
			return nil, fmt.Errorf("not enough arguments")
		}
		if os.Args[3] != model.TickData && os.Args[3] != model.KehaiData {
			return nil, fmt.Errorf("invalid data type")
		}
		request = model.RequestGroupID{
			Kei:      os.Args[2],
			DataType: os.Args[3],
			GroupID:  os.Args[4],
			Command:  os.Args[5],
		}
	case model.RequestTypeToiawase:
		if len(os.Args) != 3 {
			return nil, fmt.Errorf("not enough arguments")
		}
		request = model.RequestToiawase{
			Command: os.Args[2],
		}
	default:
		return nil, fmt.Errorf("invalid request type")
	}

	config, err := configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return nil, fmt.Errorf("file not found environment_variables/environment_variables.json")
	}
	config.TickSystem.Request = request
	config.TickSystem.S3Region = region
	config.TickSystem.S3Bucket = bucket
	config.TickSystem.RequestType = requestType

	return config, nil
}
