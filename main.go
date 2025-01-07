package main

import (
	"fmt"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/gocarina/gocsv"
	"os"
)

func main() {
	anobiiFile := "examples/exported-anobii.csv"
	clientsFile, err := os.OpenFile(anobiiFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	anobiiBooks := []*anobii.Anobii{}
	if err := gocsv.UnmarshalFile(clientsFile, &anobiiBooks); err != nil {
		panic(err)
	}
	for _, client := range anobiiBooks {
		fmt.Println(client)
	}
}
