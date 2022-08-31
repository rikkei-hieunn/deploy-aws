package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"start-jushin/api"
	"start-jushin/configs"
	"start-jushin/infrastructure"
	"start-jushin/model"
)

var (
	config    *configs.Server
	groupID   string
	dataType  string
	groupLine string
	keiType   string
	startType string
)

func main() {
	err := initConfig()
	if err != nil {
		log.Error().Msgf("Load config error: %s", err.Error())

		return
	}

	err = config.TickSystem.Validate()
	if err != nil {
		log.Error().Msgf("Load config error: %s", err.Error())

		return
	}

	infra, err := infrastructure.Init(config)
	if err != nil {
		log.Error().Msgf("Error while starting instance: %s", err.Error())

		return
	}

	server := api.New(infra, config)
	ctx := context.Background()
	err = server.Start(ctx, startType, keiType, groupID, dataType, groupLine)
	if err != nil {
		log.Error().Msgf("Error while starting instance: %s", err.Error())

		return
	}
	log.Info().Msg("Command executed successfully")
}

// initConfig get configuration file from S3 and parse to object
func initConfig() error {
	var err error
	if len(os.Args) < 3 {
		return fmt.Errorf("missing arguments")
	}
	startType = os.Args[1]
	switch startType {
	case model.TypeRunAll:
		keiType = os.Args[2]
	case model.TypeRunSSS:
		if len(os.Args) != 5 {
			return fmt.Errorf("missing arguments")
		}
		keiType = os.Args[2]
		dataType = os.Args[3]
		groupID = os.Args[4]
	case model.TypeRunByGroupLine:
		if len(os.Args) != 4 {
			return fmt.Errorf("missing arguments")
		}
		keiType = os.Args[2]
		groupLine = os.Args[3]
	default:
		return fmt.Errorf("invalid request type")
	}
	if keiType != model.FirstKei && keiType != model.SecondKei {
		return fmt.Errorf("invalid kei ")
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
	config.S3Bucket = bucket
	config.S3Region = region

	return nil
}
