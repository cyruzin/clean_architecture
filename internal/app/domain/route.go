package domain

import "context"

// Route is struct used for route type.
type Route struct {
	Departure   string `json:"departure"`
	Destination string `json:"destination"`
	Price       int    `json:"price"`
}

// RouteRepository provides access to route repository.
type RouteRepository interface {
	Find(context.Context, *Route) (*Route, error)
	Create(context.Context, *Route) error
	CheckBestRoute(filePath string, query *Route) (*Route, error)
}

// RouteService provides route operations.
type RouteService interface {
	Find(ctx context.Context, query *Route) (*Route, error)
	Create(ctx context.Context, route *Route) error
	CheckBestRoute(filePath string, query *Route) (*Route, error)
}
