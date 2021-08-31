package contextcheck_test

import (
	"log"
	"testing"

	"github.com/sylvia7788/contextcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, contextcheck.NewAnalyzer(), "a")
}
