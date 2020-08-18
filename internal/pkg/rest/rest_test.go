package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	csvstorage "github.com/cyruzin/bexs_challenge/internal/app/storage/file"

	"github.com/cyruzin/bexs_challenge/internal/app/http/controller"

	routeentity "github.com/cyruzin/bexs_challenge/internal/app/entity"

	"github.com/cyruzin/bexs_challenge/internal/pkg/rest"
)

func TestToJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/route", nil)
	if err != nil {
		t.Fatal(err)
	}

	csvRepository := csvstorage.NewCSVRepository()
	routeHandlers := controller.NewHandler(csvRepository)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routeHandlers.Show)

	handler.ServeHTTP(rr, req)

	rest.ToJSON(rr, http.StatusOK, &routeentity.Route{})
}
