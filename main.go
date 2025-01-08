package main

import (
	"fmt"
	"github.com/bu3/anobii-to-goodreads/mapping"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"os"
)

func main() {
	anobiiFile := "examples/exported-anobii.csv"
	clientsFile, err := os.OpenFile(anobiiFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	anobiiBooks, err := anobii.Read(anobiiFile)
	if err != nil {
		panic(err)
	}
	for _, client := range anobiiBooks {
		fmt.Println(client)
	}

	mapper := mapping.AnobiiToGoodReadsMapper{}
	mapper.MapList(anobiiBooks)
}
