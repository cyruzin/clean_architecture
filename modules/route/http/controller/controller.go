package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/clean_architecture/entities"
	"github.com/cyruzin/clean_architecture/pkg/rest"
	"github.com/go-chi/chi"
)

// RouteHandler is struct that implements Route Entity.
type RouteHandler struct {
	routeUseCase entities.RouteUseCase
}

// NewHandler will instantiate the handlers.
func NewHandler(c *chi.Mux, r entities.RouteUseCase) {
	handler := RouteHandler{routeUseCase: r}

	c.Route("/route", func(r chi.Router) {
		r.Get("/", handler.Find)
		r.Post("/", handler.Create)
	})
}

// Find finds the best route.
func (h *RouteHandler) Find(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if params["departure"] == nil ||
		params["destination"] == nil {
		rest.InvalidRequest(w, r, entities.ErrParams, entities.ErrParams.Error(), http.StatusBadRequest)
		return
	}

	query := &entities.Route{
		Departure:   params["departure"][0],
		Destination: params["destination"][0],
		Price:       0,
	}

	route, err := h.routeUseCase.Find(r.Context(), query)
	if err != nil {
		rest.InvalidRequest(w, r, err, err.Error(), http.StatusNotFound)
		return
	}

	rest.ToJSON(w, http.StatusOK, &route)
}

// Create creates new routes.
func (h *RouteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var route entities.Route

	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		rest.InvalidRequest(w, r, err, entities.ErrCreate.Error(), http.StatusUnprocessableEntity)
		return
	}

	if route.Departure == "" ||
		route.Destination == "" ||
		route.Price <= 0 {
		rest.InvalidRequest(w, r, err, entities.ErrFields.Error(), http.StatusBadRequest)
		return
	}

	err = h.routeUseCase.Create(r.Context(), &route)
	if err != nil {
		rest.InvalidRequest(w, r, err, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	rest.ToJSON(
		w,
		http.StatusOK,
		&rest.APIMessage{Message: "Route created", Status: http.StatusCreated},
	)
}
