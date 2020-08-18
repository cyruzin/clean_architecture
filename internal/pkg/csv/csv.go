package csv

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"

	"github.com/rs/zerolog/log"
)

// read reads the content of the CSV file.
func read(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open the CSV file")
		return nil, err
	}

	r := csv.NewReader(csvFile)

	var parsedFile [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error().Err(err).Msg("Failed to read the CSV file")
			return [][]string{}, err
		}

		parsedFile = append(parsedFile, record)
	}

	return parsedFile, err
}

// Write writes the new content to the end of the file.
func Write(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Could not open CSV file to write")
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		log.Error().Err(err).Msg("Failed to write to the CSV file")
		return err
	}

	return nil
}

// parse parses the CSV file into a slice of Route.
func parse(filePath string) ([]routeentity.Route, error) {
	rawCSV, err := read(filePath)
	if err != nil {
		return nil, err
	}

	var parsedCSV []routeentity.Route

	for _, column := range rawCSV {

		price, err := strconv.Atoi(column[2])
		if err != nil {
			return nil, err
		}

		route := routeentity.Route{
			Departure:   column[0],
			Destination: column[1],
			Price:       price,
		}

		parsedCSV = append(parsedCSV, route)
	}

	return parsedCSV, nil
}

// CheckBestRoute iterate over the slice of Route and returns
// a string with the best route.
func CheckBestRoute(
	filePath string,
	query routeentity.Route,
) (routeentity.Route, error) {
	routes, err := parse(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse the CSV file")
		return routeentity.Route{}, err
	}

	var filteredRoutes []routeentity.Route

	for _, route := range routes {
		if query.Departure == route.Departure &&
			query.Destination == route.Destination {
			filteredRoutes = append(filteredRoutes, route)
		}
	}

	sort.SliceStable(filteredRoutes, func(i, j int) bool {
		return filteredRoutes[i].Price < filteredRoutes[j].Price
	})

	bestRoute := routeentity.Route{
		Departure:   query.Departure,
		Destination: query.Destination,
		Price:       filteredRoutes[0].Price,
	}

	return bestRoute, nil
}

// CheckDuplicateRoute checks if the given route is already on the file.
func CheckDuplicateRoute(
	filePath string,
	query *routeentity.Route,
) (bool, error) {
	routes, err := parse(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse the CSV file")
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
