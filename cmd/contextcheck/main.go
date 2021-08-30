package main

import (
	"github.com/1227977886/contextcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(contextcheck.NewAnalyzer())
}
