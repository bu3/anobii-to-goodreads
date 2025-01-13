package convert

import (
	"fmt"
	"github.com/bu3/anobii-to-goodreads/file"
	"github.com/bu3/anobii-to-goodreads/mapping"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
)

type Converter struct {
	anobiiFileReader anobii.AnobiiFileReader
}

func Convert(inputFile string, outputFile string) error {
	fileManager := file.New()
	anobiiFileReader := anobii.New(fileManager)
	goodReadsFile := goodreads.New(fileManager)
	anobiiBooks, err := anobiiFileReader.Read(inputFile)
	if err != nil {
		panic(err)
	}

	mapper := mapping.AnobiiToGoodReadsMapper{}
	goodReadsItems, err := mapper.MapList(anobiiBooks)
	if err != nil {
		panic(err)
	}
	for _, goodReadsItem := range *goodReadsItems {
		fmt.Println(goodReadsItem)
	}

	goodReadsFile.Write(outputFile, goodReadsItems)

	return nil
}
