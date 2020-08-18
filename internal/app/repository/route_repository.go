package routerepository

import (
	"context"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"
)

// RouteRepository provides access to route repository.
type RouteRepository interface {
	Find(context.Context, routeentity.Route) (*routeentity.Route, error)
	Create(context.Context, *routeentity.Route) error
}
