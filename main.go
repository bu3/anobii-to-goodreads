package main

import (
	"github.com/bu3/anobii-to-goodreads/convert"
	"github.com/bu3/anobii-to-goodreads/mapping"
	"os"
)

func main() {
	anobiiFile := "examples/exported-anobii.csv"
	mapper := mapping.New()
	converter := convert.Converter{
		Mapper: mapper,
	}
	_ = converter.Convert(anobiiFile, os.Args[1])
}
