package csv_test

import (
	"context"
	"testing"
	"time"

	"github.com/cyruzin/clean_architecture/domain"
	"github.com/cyruzin/clean_architecture/modules/route/repository/csv"
)

var repo = csv.NewCSVRepository()

func TestFind(t *testing.T) {
	route := &domain.Route{
		Departure:   "RJ",
		Destination: "SP",
	}

	_, err := repo.Find(context.TODO(), route)
	if err != nil {
		t.Error(err)
	}
}

func TestFindFail(t *testing.T) {
	route := &domain.Route{
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
	route := &domain.Route{
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
	route := &domain.Route{
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
