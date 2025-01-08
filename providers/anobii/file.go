package anobii

import (
	"github.com/gocarina/gocsv"
	"os"
)

func Read(anobiiFile string) ([]*Anobii, error) {
	clientsFile, err := os.OpenFile(anobiiFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer clientsFile.Close()

	var anobiiBooks []*Anobii
	if err := gocsv.UnmarshalFile(clientsFile, &anobiiBooks); err != nil {
		return nil, err
	}

	return anobiiBooks, nil
}
