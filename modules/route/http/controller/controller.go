package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyruzin/clean_architecture/domain"
	"github.com/cyruzin/clean_architecture/pkg/rest"
	"github.com/go-chi/chi"
)

// RouteHandler is struct that implements Route Service.
type RouteHandler struct {
	routeService domain.RouteService
}

// NewHandler will instantiate the handlers.
func NewHandler(c *chi.Mux, rs domain.RouteService) {
	handler := RouteHandler{routeService: rs}

	c.Route("/route", func(r chi.Router) {
		r.Get("/", handler.Find)
		r.Post("/", handler.Create)
	})
}

// Find finds the best route.
func (h *RouteHandler) Find(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	fmt.Println(params)

	if params["departure"] == nil ||
		params["destination"] == nil {
		rest.InvalidRequest(w, r, domain.ErrParams, domain.ErrParams.Error(), http.StatusBadRequest)
		return
	}

	query := &domain.Route{
		Departure:   params["departure"][0],
		Destination: params["destination"][0],
		Price:       0,
	}

	route, err := h.routeService.Find(r.Context(), query)
	if err != nil {
		rest.InvalidRequest(w, r, err, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	rest.ToJSON(w, http.StatusOK, &route)
}

// Create creates new routes.
func (h *RouteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var route domain.Route

	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		rest.InvalidRequest(w, r, err, domain.ErrCreate.Error(), http.StatusUnprocessableEntity)
		return
	}

	if route.Departure == "" ||
		route.Destination == "" ||
		route.Price <= 0 {
		rest.InvalidRequest(w, r, err, domain.ErrFields.Error(), http.StatusBadRequest)
		return
	}

	err = h.routeService.Create(r.Context(), &route)
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
