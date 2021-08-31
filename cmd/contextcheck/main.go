package main

import (
	"github.com/sylvia7788/contextcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(contextcheck.NewAnalyzer())
}
