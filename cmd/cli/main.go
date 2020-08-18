package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"
	"github.com/cyruzin/bexs_challenge/internal/pkg/csv"
)

const filePath = "../../assets/routes.csv"

func main() {
	args := os.Args[1:]

	query := routeentity.Route{
		Departure:   args[0],
		Destination: args[1],
		Price:       0,
	}

	route, err := csv.CheckBestRoute(filePath, query)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return
	}

	bestRoute := fmt.Sprintf(
		"best route: %s - %s > $%d",
		route.Departure,
		route.Destination,
		route.Price,
	)

	log.Info().Msg(bestRoute)
}
