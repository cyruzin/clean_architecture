package csv_test

import (
	"context"
	"testing"
	"time"

	"github.com/cyruzin/clean_architecture/entities"
	"github.com/cyruzin/clean_architecture/modules/route/presenter/csv"
	"github.com/cyruzin/clean_architecture/pkg/util"
)

var repo = csv.NewCSVPresenter(util.PathBuilder("/assets/routes.csv"))

func TestFind(t *testing.T) {
	route := &entities.Route{
		Departure:   "RJ",
		Destination: "SP",
	}

	_, err := repo.Find(context.TODO(), route)
	if err != nil {
		t.Error(err)
	}
}

func TestFindFail(t *testing.T) {
	route := &entities.Route{
		Departure:   "RJ1",
		Destination: "SP2",
	}

	_, err := repo.Find(context.TODO(), route)
	if err == nil {
		t.Error(err)
	}

	route.Departure = ""
	route.Destination = ""

	_, err = repo.Find(context.TODO(), route)
	if err == nil {
		t.Error(err)
	}
}

func TestCreate(t *testing.T) {
	route := &entities.Route{
		Departure:   "RJ",
		Destination: "SP",
		Price:       int(time.Now().UnixNano()),
	}

	err := repo.Create(context.TODO(), route)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateFail(t *testing.T) {
	route := &entities.Route{
		Departure:   "",
		Destination: "",
		Price:       0,
	}

	err := repo.Create(context.TODO(), route)
	if err == nil {
		t.Error(err)
	}

	route.Departure = "GRU"
	route.Destination = "BRC"
	route.Price = 10

	err = repo.Create(context.TODO(), route)
	if err == nil {
		t.Error(err)
	}
}
