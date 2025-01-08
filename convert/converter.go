package convert

import (
	"fmt"
	"github.com/bu3/anobii-to-goodreads/mapping"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
	"os"
)

func Convert(inputFile string, outputFile string) error {
	clientsFile, err := os.OpenFile(inputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	anobiiBooks, err := anobii.Read(inputFile)
	if err != nil {
		panic(err)
	}
	for _, anobiiBook := range anobiiBooks {
		fmt.Println(anobiiBook)
	}

	mapper := mapping.AnobiiToGoodReadsMapper{}
	goodReadsItems, err := mapper.MapList(anobiiBooks)
	if err != nil {
		panic(err)
	}
	for _, goodReadsItem := range *goodReadsItems {
		fmt.Println(goodReadsItem)
	}
	goodreads.Write(outputFile, goodReadsItems)

	return nil
}
