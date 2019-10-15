package defaultiota_test

import (
	"testing"

	"github.com/gostaticanalysis/ulinter/defaultiota"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestInvalidSuffix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, defaultiota.Analyzer, "a")
}

func TestValidSuffix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, defaultiota.Analyzer, "b")
}
