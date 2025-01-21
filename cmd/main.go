package main

import (
	"flag"
	"github.com/bu3/anobii-to-goodreads/pkg/convert"
)

var (
	inputFile  string
	outputFile string
)

func main() {
	flag.StringVar(&inputFile, "in", "", "path to the input file")
	flag.StringVar(&outputFile, "out", "", "path to the converted file")
	flag.Parse()

	if inputFile == "" || outputFile == "" {
		panic("input and output file must be specified")
	}

	err := convert.Convert(inputFile, outputFile)
	if err != nil {
		panic("Error converting input file to output file: " + err.Error())
	}
}
