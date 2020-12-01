package controller

import (
	"encoding/json"
	"net/http"

	routeentity "github.com/cyruzin/clean_architecture/internal/app/entity/route"
	"github.com/cyruzin/clean_architecture/internal/pkg/rest"

	routeservice "github.com/cyruzin/clean_architecture/internal/app/service/route"
)

const (
	errParams = "Departure/Destination param missing"
	errFields = "All fields are required"
	errCreate = "Could not create the route"
)

// RouteHandler interface for the route handlers.
type RouteHandler interface {
	Show(http.ResponseWriter, *http.Request)
	Store(http.ResponseWriter, *http.Request)
}

type csvHandler struct {
	routeService routeservice.RouteService
}

// NewHandler will instantiate the handlers.
func NewHandler(routeService routeservice.RouteService) RouteHandler {
	return &csvHandler{routeService}
}

// Show shows the best route.
func (h *csvHandler) Show(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if params["departure"] == nil ||
		params["destination"] == nil {
		rest.InvalidRequest(w, r, nil, errParams, http.StatusBadRequest)
		return
	}

	query := routeentity.Route{
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

// Store stores new routes.
func (h *csvHandler) Store(w http.ResponseWriter, r *http.Request) {
	var route routeentity.Route

	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		rest.InvalidRequest(w, r, err, errCreate, http.StatusUnprocessableEntity)
		return
	}

	if route.Departure == "" ||
		route.Destination == "" ||
		route.Price <= 0 {
		rest.InvalidRequest(w, r, nil, errFields, http.StatusBadRequest)
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
