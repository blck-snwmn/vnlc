package main

import (
	"github/blck-snwmn/vnlc"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(vnlc.Analyzer)
}
