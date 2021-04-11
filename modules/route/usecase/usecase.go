package usecase

import (
	"context"

	"github.com/cyruzin/clean_architecture/entities"
)

type routeUseCase struct {
	routePresenter entities.RoutePresenter
}

func NewRouteUseCase(r entities.RoutePresenter) entities.RouteUseCase {
	return &routeUseCase{r}
}

// Find finds the best route.
func (s *routeUseCase) Find(ctx context.Context, r *entities.Route) (*entities.Route, error) {
	route, err := s.routePresenter.Find(ctx, r)
	if err != nil {
		return nil, err
	}

	return route, nil
}

// Create adds the new route to the csv file.
func (s *routeUseCase) Create(ctx context.Context, route *entities.Route) error {
	if err := s.routePresenter.Create(ctx, route); err != nil {
		return err
	}

	return nil
}

// CheckBestRoute checks the best route.
func (s *routeUseCase) CheckBestRoute(r *entities.Route) (*entities.Route, error) {
	route, err := s.routePresenter.CheckBestRoute(r)
	if err != nil {
		return nil, err
	}

	return route, nil
}
