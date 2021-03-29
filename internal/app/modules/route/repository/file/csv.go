package storage

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cyruzin/clean_architecture/internal/app/domain"
	"github.com/cyruzin/clean_architecture/pkg/csv"
	"github.com/cyruzin/clean_architecture/pkg/util"
	"github.com/rs/zerolog/log"
)

var filePath = util.PathBuilder("./assets/routes.csv")

type csvRepository struct{}

// NewCSVRepository access the routes repository.
func NewCSVRepository() domain.RouteRepository {
	return &csvRepository{}
}

// Find finds the best route.
func (c *csvRepository) Find(ctx context.Context, query *domain.Route) (*domain.Route, error) {
	bestRoute, err := c.CheckBestRoute(filePath, query)
	if err != nil {
		return nil, err
	}

	return bestRoute, nil
}

// Create creates a new route.
func (c *csvRepository) Create(ctx context.Context, route *domain.Route) error {
	duplicate, err := c.checkDuplicateRoute(filePath, route)
	if err != nil {
		return err
	}

	if duplicate {
		return domain.ErrConflict
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

// CheckBestRoute iterate over the slice of Route and returns
// a string with the best route.
func (c *csvRepository) CheckBestRoute(filePath string, query *domain.Route) (*domain.Route, error) {
	routes, err := c.parse(filePath)
	if err != nil {
		log.Error().Err(err).Stack().Msg(err.Error())
		return &domain.Route{}, err
	}

	var filteredRoutes []domain.Route

	for _, route := range routes {
		if query.Departure == route.Departure &&
			query.Destination == route.Destination {
			filteredRoutes = append(filteredRoutes, route)
		}
	}

	sort.SliceStable(filteredRoutes, func(i, j int) bool {
		return filteredRoutes[i].Price < filteredRoutes[j].Price
	})

	if len(filteredRoutes) == 0 {
		return &domain.Route{}, domain.ErrNotFound
	}

	bestRoute := domain.Route{
		Departure:   query.Departure,
		Destination: query.Destination,
		Price:       filteredRoutes[0].Price,
	}

	return &bestRoute, nil
}

// checkDuplicateRoute checks if the given route is already on the file.
func (c *csvRepository) checkDuplicateRoute(
	filePath string,
	query *domain.Route,
) (bool, error) {
	routes, err := c.parse(filePath)
	if err != nil {
		log.Error().Err(err).Stack().Msg(err.Error())
		return false, err
	}

	for _, route := range routes {
		if query.Departure == route.Departure &&
			query.Destination == route.Destination &&
			query.Price == route.Price {
			return true, nil
		}
	}

	return false, nil
}

// parse parses the CSV file into a slice of Route.
func (c *csvRepository) parse(filePath string) ([]domain.Route, error) {
	rawCSV, err := csv.Read(filePath)
	if err != nil {
		return nil, err
	}

	var parsedCSV []domain.Route

	for _, column := range rawCSV {

		price, err := strconv.Atoi(column[2])
		if err != nil {
			return nil, err
		}

		route := domain.Route{
			Departure:   column[0],
			Destination: column[1],
			Price:       price,
		}

		parsedCSV = append(parsedCSV, route)
	}

	return parsedCSV, nil
}
