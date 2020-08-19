package main

import (
	"fmt"
	"log"
	"os"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity/route"
	"github.com/cyruzin/bexs_challenge/internal/pkg/csv"
)

const filePath = "../../assets/routes.csv"

func usage() {
	log.Println(
		`Wrong usage, check the examples below:

Development example: go run main.go BBB AAA
Production example: ./cli BBB AAA`)
}

func main() {
	log.SetFlags(0)

	if len(os.Args[1:]) < 2 {
		usage()
		return
	}

	args := os.Args[1:]

	query := routeentity.Route{
		Departure:   args[0],
		Destination: args[1],
		Price:       0,
	}

	route, err := csv.CheckBestRoute(filePath, query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	bestRoute := fmt.Sprintf(
		"best route: %s - %s > $%d",
		route.Departure,
		route.Destination,
		route.Price,
	)

	log.Println(bestRoute)
}
