package main

import (
	"context"
	"net/http"
	"os/signal"
	"time"

	"os"

	"github.com/cyruzin/clean_architecture/modules/route/http/controller"
	routeMiddleware "github.com/cyruzin/clean_architecture/modules/route/http/middleware"
	storage "github.com/cyruzin/clean_architecture/modules/route/repository/file"
	"github.com/cyruzin/clean_architecture/modules/route/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/cyruzin/clean_architecture/config"
	"github.com/go-chi/chi/middleware"
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

	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(routeMiddleware.LoggerMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Clean Architecture example"))
	})

	controller.NewHandler(router, routeService)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		ReadTimeout:       cfg.ReadTimeOut,
		ReadHeaderTimeout: cfg.ReadHeaderTimeOut,
		WriteTimeout:      cfg.WriteTimeOut,
		IdleTimeout:       cfg.IdleTimeOut,
		Handler:           router,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		<-gracefulStop

		log.Info().Msg("shutting down the server...")
		if err := srv.Shutdown(ctx); err != nil {
			log.Error().Err(err).Stack().Msg("server failed to shutdown")
		}
		close(idleConnsClosed)
	}()

	log.Info().Msgf("listening on port: %s", cfg.Port)
	log.Info().Msg("you're good to go! :)")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Error().Err(err).Stack().Msg("server failed to start")
	}

	<-idleConnsClosed
}
