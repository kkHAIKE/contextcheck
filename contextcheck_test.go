package contextcheck_test

import (
	"log"
	"testing"

	"github.com/1227977886/contextcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, contextcheck.NewAnalyzer(), "a")
}
