package entities

import "context"

// Route is struct used for route type.
type Route struct {
	Departure   string `json:"departure"`
	Destination string `json:"destination"`
	Price       int    `json:"price"`
}

// RoutePresenter provides access to route presenter.
type RoutePresenter interface {
	Find(context.Context, *Route) (*Route, error)
	Create(context.Context, *Route) error
	CheckBestRoute(r *Route) (*Route, error)
}

// RouteUseCase provides route operations.
type RouteUseCase interface {
	Find(ctx context.Context, r *Route) (*Route, error)
	Create(ctx context.Context, route *Route) error
	CheckBestRoute(r *Route) (*Route, error)
}
