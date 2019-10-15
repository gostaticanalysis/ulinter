package main

import (
	"github.com/gostaticanalysis/ulinter/defaultiota"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(defaultiota.Analyzer) }
