package csvstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/cyruzin/bexs_challenge/internal/pkg/csv"

	routerepository "github.com/cyruzin/bexs_challenge/internal/app/repository"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"
)

const filePath = "../../assets/routes.csv"

type csvRepository struct{}

// NewCSVRepository access the routes repository.
func NewCSVRepository() routerepository.RouteRepository {
	return &csvRepository{}
}

// Find finds the best route.
func (c *csvRepository) Find(
	ctx context.Context,
	query routeentity.Route,
) (*routeentity.Route, error) {
	bestRoute, err := csv.CheckBestRoute(filePath, query)
	if err != nil {
		return nil, err
	}

	return &bestRoute, nil
}

// Create creates a new route.
func (c *csvRepository) Create(
	ctx context.Context,
	route *routeentity.Route,
) error {
	duplicate, err := csv.CheckDuplicateRoute(filePath, route)
	if err != nil {
		return err
	}

	if duplicate {
		return errors.New("That route already exists")
	}

	newEntry := fmt.Sprintf(
		"%s,%s,%d",
		route.Departure,
		route.Destination,
		route.Price,
	)

	if err := csv.Write(filePath, newEntry); err != nil {
		return err
	}

	return nil
}
