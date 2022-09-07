package main

import (
	"github.com/kkHAIKE/contextcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(contextcheck.NewAnalyzer(contextcheck.Configuration{}))
}
