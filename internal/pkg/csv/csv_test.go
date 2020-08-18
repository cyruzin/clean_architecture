package csv_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"
	"github.com/cyruzin/bexs_challenge/internal/pkg/csv"
)

const filePath = "../../../assets/routes.csv"

func TestCheckBestRouteSuccess(t *testing.T) {
	query := routeentity.Route{
		Departure:   "BBB",
		Destination: "AAA",
		Price:       0,
	}

	bestRoute, err := csv.CheckBestRoute(filePath, query)
	if err != nil {
		t.Error("failed to check the best route")
	}

	if bestRoute.Departure == "" {
		t.Error("empty departure")
	}
}

func TestCheckBestRouteFail(t *testing.T) {
	query := routeentity.Route{
		Departure:   "",
		Destination: "AAA",
		Price:       0,
	}

	_, err := csv.CheckBestRoute(filePath, query)
	if err == nil {
		t.Error("failed to check empty params")
	}
}

func TestDuplicatRouteSuccess(t *testing.T) {
	query := routeentity.Route{
		Departure:   "BBB",
		Destination: "AAA",
		Price:       5,
	}

	_, err := csv.CheckDuplicateRoute(filePath, &query)
	if err != nil {
		t.Error("failed to check duplicate route")
	}
}

func TestDuplicateRouteFail(t *testing.T) {
	query := routeentity.Route{
		Departure:   "BBB",
		Destination: "AAA",
		Price:       99,
	}

	_, err := csv.CheckDuplicateRoute(filePath, &query)
	if err != nil {
		t.Error("failed to check duplicate route")
	}
}

func TestWriteSuccess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	min := 500
	max := 100000

	randomPrice := rand.Intn(max-min+1) + min

	strPrice := strconv.Itoa(randomPrice)

	testRoute := fmt.Sprintf("YYY,ZZZ,%s", strPrice)

	err := csv.Write(filePath, testRoute)
	if err != nil {
		t.Error("failed to write to the csv file")
	}
}

func TestWriteFail(t *testing.T) {
	err := csv.Write(filePath, "")
	if err == nil {
		t.Errorf("failed to prevent duplicate write: %s", err.Error())
	}
}
