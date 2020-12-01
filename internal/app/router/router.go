package router

import (
	"net/http"
	"time"

	"github.com/cyruzin/clean_architecture/internal/app/http/controller"
	"github.com/cyruzin/clean_architecture/internal/app/http/middleware"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// New initiates all routes.
func New(h controller.RouteHandler) http.Handler {
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
	router.Use(chiMiddleware.Timeout(60 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.LoggerMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Clear Architecture by Cyro Dubeux"))
	})

	router.Route("/route", func(router chi.Router) {
		router.Get("/", h.Show)
		router.Post("/", h.Store)
	})

	return router
}
