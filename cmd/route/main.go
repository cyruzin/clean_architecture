package main

import (
	"context"

	"os"

	"github.com/cyruzin/clean_architecture/internal/app/modules/route/http/controller"
	storage "github.com/cyruzin/clean_architecture/internal/app/modules/route/repository/file"
	"github.com/cyruzin/clean_architecture/internal/app/modules/route/service"
	"github.com/cyruzin/clean_architecture/internal/app/server"

	"github.com/cyruzin/clean_architecture/internal/app/router"

	"github.com/cyruzin/clean_architecture/internal/app/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if cfg.EnvMode == "production" {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Debug().Msg("running in production mode")
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Debug().Msg("running in development mode")
	}

	routeRepository := storage.NewCSVRepository()
	routeService := service.NewService(routeRepository)

	routeHandlers := controller.NewHandler(routeService)
	routes := router.New(routeHandlers)

	server.Start(ctx, cfg, routes)
}
