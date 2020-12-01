package routerepository

import (
	"context"

	routeentity "github.com/cyruzin/clean_architecture/internal/app/entity/route"
)

// RouteRepository provides access to route repository.
type RouteRepository interface {
	Find(context.Context, routeentity.Route) (*routeentity.Route, error)
	Create(context.Context, *routeentity.Route) error
}
