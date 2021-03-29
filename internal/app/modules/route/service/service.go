package service

import (
	"context"

	"github.com/cyruzin/clean_architecture/internal/app/domain"
)

type routeService struct {
	routeRepository domain.RouteRepository
}

// NewService creates a service with the necessary dependencies.
func NewService(r domain.RouteRepository) domain.RouteService {
	return &routeService{r}
}

// Find finds the best route.
func (s *routeService) Find(ctx context.Context, query *domain.Route) (*domain.Route, error) {
	route, err := s.routeRepository.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	return route, nil
}

// Create adds the new route to the csv file.
func (s *routeService) Create(ctx context.Context, route *domain.Route) error {
	if err := s.routeRepository.Create(ctx, route); err != nil {
		return err
	}

	return nil
}

// CheckBestRoute checks the best route.
func (s *routeService) CheckBestRoute(filePath string, query *domain.Route) (*domain.Route, error) {
	route, err := s.routeRepository.CheckBestRoute(filePath, query)
	if err != nil {
		return nil, err
	}

	return route, nil
}
