package numstr_test

import (
	"testing"

	"github.com/gostaticanalysis/ulinter/numstr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestOrdiaryPattern(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, numstr.Analyzer, "a")
}

func TestNamedPackage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, numstr.Analyzer, "b")
}
