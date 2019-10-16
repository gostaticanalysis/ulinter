package numstr_test

import (
	"testing"

	"github.com/gostaticanalysis/ulinter/numstr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, numstr.Analyzer, "a")
}