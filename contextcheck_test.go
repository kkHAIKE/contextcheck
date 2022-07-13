package contextcheck_test

import (
	"log"
	"testing"

	"github.com/sylvia7788/contextcheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/packages"
)

func Test(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	testdata := analysistest.TestData()
	analyzer := contextcheck.NewAnalyzer()
	analyzer.Run = contextcheck.NewRun([]*packages.Package{{PkgPath: "a"}})
	analysistest.Run(t, testdata, analyzer, "a")
}
