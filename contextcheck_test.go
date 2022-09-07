package contextcheck_test

import (
	"log"
	"testing"

	"github.com/kkHAIKE/contextcheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/packages"
)

func Test(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	testdata := analysistest.TestData()
	analyzer := contextcheck.NewAnalyzer(contextcheck.Configuration{})
	analyzer.Run = contextcheck.NewRun([]*packages.Package{{PkgPath: "a"}}, false)
	analysistest.Run(t, testdata, analyzer, "a")
}
