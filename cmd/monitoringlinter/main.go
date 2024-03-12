package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/assafad/monitoringlinter"
)

func main() {
	singlechecker.Main(monitoringlinter.NewAnalyzer())
}
