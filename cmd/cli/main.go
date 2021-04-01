package main

import (
	"log"
	"os"

	"github.com/cyruzin/clean_architecture/domain"
	storage "github.com/cyruzin/clean_architecture/modules/route/repository/file"
	"github.com/cyruzin/clean_architecture/modules/route/service"
	"github.com/cyruzin/clean_architecture/pkg/util"
)

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

	query := &domain.Route{
		Departure:   args[0],
		Destination: args[1],
		Price:       0,
	}

	routeRepository := storage.NewCSVRepository()
	routeService := service.NewService(routeRepository)

	route, err := routeService.CheckBestRoute(util.PathBuilder("/assets/routes.csv"), query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf(
		"the best route from %s to %s costs $%d",
		route.Departure,
		route.Destination,
		route.Price,
	)
}
