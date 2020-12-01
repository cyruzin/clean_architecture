package main

import (
	"context"

	"github.com/cyruzin/clean_architecture/internal/app/server"

	"github.com/cyruzin/clean_architecture/internal/app/http/controller"

	"github.com/cyruzin/clean_architecture/internal/app/router"
	csvstorage "github.com/cyruzin/clean_architecture/internal/app/storage/file"

	"github.com/cyruzin/clean_architecture/internal/app/config"
	"github.com/rs/zerolog"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	csvRepository := csvstorage.NewCSVRepository()
	routeHandlers := controller.NewHandler(csvRepository)
	routes := router.New(routeHandlers)

	server.Start(ctx, cfg, routes)
}
