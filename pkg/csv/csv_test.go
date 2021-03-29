package csv_test

import (
	"testing"

	"github.com/cyruzin/clean_architecture/pkg/csv"
	"github.com/cyruzin/clean_architecture/pkg/util"
)

var filePath = util.PathBuilder("/assets/routes.csv")

func TestRead(t *testing.T) {
	_, err := csv.Read(filePath)
	if err != nil {
		t.Error(err)
	}
}

func TestWrite(t *testing.T) {
	if err := csv.Write(filePath, "RJ,SP,5"); err != nil {
		t.Error(err)
	}
}
