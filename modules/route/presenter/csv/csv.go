package csv

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cyruzin/clean_architecture/entities"
	"github.com/cyruzin/clean_architecture/pkg/csv"
	"github.com/rs/zerolog/log"
)

type csvPresenter struct {
	filePath string
}

func NewCSVPresenter(fp string) entities.RoutePresenter {
	return &csvPresenter{filePath: fp}
}

// Find finds the best route.
func (c *csvPresenter) Find(ctx context.Context, r *entities.Route) (*entities.Route, error) {
	if r.Departure == "" ||
		r.Destination == "" {
		return nil, entities.ErrParams
	}

	bestRoute, err := c.CheckBestRoute(r)
	if err != nil {
		return nil, err
	}

	return bestRoute, nil
}

// Create creates a new route.
func (c *csvPresenter) Create(ctx context.Context, r *entities.Route) error {
	if r.Departure == "" ||
		r.Destination == "" ||
		r.Price <= 0 {
		return entities.ErrFields
	}

	duplicate, err := c.checkDuplicateRoute(r)
	if err != nil {
		return err
	}

	if duplicate {
		return entities.ErrConflict
	}

	newEntry := fmt.Sprintf(
		"%s,%s,%d",
		r.Departure,
		r.Destination,
		r.Price,
	)

	if err := csv.Write(c.filePath, newEntry); err != nil {
		return err
	}

	return nil
}

// CheckBestRoute iterate over the slice of Route and returns
// a string with the best route.
func (c *csvPresenter) CheckBestRoute(r *entities.Route) (*entities.Route, error) {
	routes, err := c.parse()
	if err != nil {
		log.Error().Err(err).Stack().Msg(err.Error())
		return &entities.Route{}, err
	}

	var filteredRoutes []entities.Route

	for _, route := range routes {
		if r.Departure == route.Departure &&
			r.Destination == route.Destination {
			filteredRoutes = append(filteredRoutes, route)
		}
	}

	sort.SliceStable(filteredRoutes, func(i, j int) bool {
		return filteredRoutes[i].Price < filteredRoutes[j].Price
	})

	if len(filteredRoutes) == 0 {
		return &entities.Route{}, entities.ErrNotFound
	}

	bestRoute := entities.Route{
		Departure:   r.Departure,
		Destination: r.Destination,
		Price:       filteredRoutes[0].Price,
	}

	return &bestRoute, nil
}

// checkDuplicateRoute checks if the given route is already on the file.
func (c *csvPresenter) checkDuplicateRoute(r *entities.Route) (bool, error) {
	routes, err := c.parse()
	if err != nil {
		log.Error().Err(err).Stack().Msg(err.Error())
		return false, err
	}

	for _, route := range routes {
		if r.Departure == route.Departure &&
			r.Destination == route.Destination &&
			r.Price == route.Price {
			return true, nil
		}
	}

	return false, nil
}

// parse parses the CSV file into a slice of Route.
func (c *csvPresenter) parse() ([]entities.Route, error) {
	rawCSV, err := csv.Read(c.filePath)
	if err != nil {
		return nil, err
	}

	var parsedCSV []entities.Route

	for _, column := range rawCSV {

		price, err := strconv.Atoi(column[2])
		if err != nil {
			return nil, err
		}

		route := entities.Route{
			Departure:   column[0],
			Destination: column[1],
			Price:       price,
		}

		parsedCSV = append(parsedCSV, route)
	}

	return parsedCSV, nil
}
