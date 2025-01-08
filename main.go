package main

import (
	"github.com/bu3/anobii-to-goodreads/convert"
	"os"
)

func main() {
	anobiiFile := "examples/exported-anobii.csv"
	convert.Convert(anobiiFile, os.Args[1])
}
