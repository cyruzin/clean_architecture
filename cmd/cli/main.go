package main

import (
	"log"
	"os"

	"github.com/cyruzin/clean_architecture/entities"
	routePresenter "github.com/cyruzin/clean_architecture/modules/route/presenter/csv"
	routeUseCase "github.com/cyruzin/clean_architecture/modules/route/usecase"
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

	r := &entities.Route{
		Departure:   args[0],
		Destination: args[1],
		Price:       0,
	}

	routePresenter := routePresenter.NewCSVPresenter(util.PathBuilder("/assets/routes.csv"))
	routeUseCase := routeUseCase.NewRouteUseCase(routePresenter)

	route, err := routeUseCase.CheckBestRoute(r)
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
