package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/cyruzin/clean_architecture/internal/app/config"
	"github.com/rs/zerolog/log"
)

// Start starts the server.
func Start(
	ctx context.Context,
	cfg *config.Config,
	routes http.Handler,
) {
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		ReadTimeout:       cfg.ReadTimeOut,
		ReadHeaderTimeout: cfg.ReadHeaderTimeOut,
		WriteTimeout:      cfg.WriteTimeOut,
		IdleTimeout:       cfg.IdleTimeOut,
		Handler:           routes,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		<-gracefulStop

		log.Info().Msg("shutting down the server...")
		if err := srv.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("server failed to shutdown")
		}
		close(idleConnsClosed)
	}()

	log.Info().Msgf("listening on port: %s", cfg.Port)
	log.Info().Msg("you're good to go! :)")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Error().Err(err).Msg("server failed to start")
	}

	<-idleConnsClosed
}
