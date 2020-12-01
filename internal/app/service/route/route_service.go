package routeservice

import (
	"context"

	routerepository "github.com/cyruzin/clean_architecture/internal/app/repository/route"

	routeentity "github.com/cyruzin/clean_architecture/internal/app/entity/route"
)

// RouteService provides route operations.
type RouteService interface {
	Find(ctx context.Context, query routeentity.Route) (*routeentity.Route, error)
	Create(ctx context.Context, route *routeentity.Route) error
}

type routeService struct {
	routeRepository routerepository.RouteRepository
}

// NewService creates a service with the necessary dependencies.
func NewService(r routerepository.RouteRepository) RouteService {
	return &routeService{r}
}

// Find finds the best route.
func (s *routeService) Find(
	ctx context.Context,
	query routeentity.Route,
) (*routeentity.Route, error) {
	route, err := s.routeRepository.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	return route, nil
}

// Create adds the new route to the csv file.
func (s *routeService) Create(ctx context.Context, route *routeentity.Route) error {
	if err := s.routeRepository.Create(ctx, route); err != nil {
		return err
	}
	return nil
}
