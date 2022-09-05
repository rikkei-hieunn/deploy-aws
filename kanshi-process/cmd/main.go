package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"kanshi-process/configs"
	handlerequest "kanshi-process/internal/handle_request"
	"os"
	"os/signal"
	"syscall"
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

	if err = config.TickSocket.Validate(); err != nil {
		log.Error().Msg(err.Error())

		return
	}

	requestHandler := handlerequest.NewRequestHandler(config)
	err = requestHandler.StartSendRequest(context.Background())
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}

	idleConnClosed := make(chan struct{})

	go func() {
		// Handle incoming exit signals from os
		// SIGTERM kill command; SIGINT interrupt by [Ctrl + C]; SIGHUP terminal hangup; SIGQUIT [Ctrl + \]
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		sig := <-sigint
		log.Info().Msgf("receive signal %s", sig.String())

		close(idleConnClosed)
	}()

	<-idleConnClosed
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
	config, err = configs.Init("environment_variables", "environment_variables.json")
	if err != nil {
		return err
	}

	return nil
}
