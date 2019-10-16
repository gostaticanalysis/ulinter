package main

import (
	"github.com/gostaticanalysis/ulinter/numstr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(numstr.Analyzer) }