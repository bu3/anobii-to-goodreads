package anobii

import (
	"github.com/bu3/anobii-to-goodreads/file"
	"github.com/gocarina/gocsv"
)

type AnobiiFileReader interface {
	Read(anobiiFile string) ([]*Anobii, error)
}

func New(manager file.Manager) AnobiiFileReader {
	return &anobiiFileReader{
		fileManager: manager,
	}
}

type anobiiFileReader struct {
	fileManager file.Manager
}

func (a *anobiiFileReader) Read(anobiiFile string) ([]*Anobii, error) {
	inputFile, err := a.fileManager.ReadFile(anobiiFile)
	if err != nil {
		return nil, err
	}
	var anobiiBooks []*Anobii
	if err := gocsv.UnmarshalBytes(inputFile, &anobiiBooks); err != nil {
		return nil, err
	}

	return anobiiBooks, nil
}
