package anobii

import (
	"github.com/bu3/anobii-to-goodreads/file"
	"github.com/gocarina/gocsv"
	"os"
)

type AnobiiFileReader interface {
	Read(anobiiFile *os.File) ([]*Anobii, error)
}

func New(manager file.Manager) AnobiiFileReader {
	return &anobiiFileReader{}
}

type anobiiFileReader struct {
	fileManager file.Manager
}

func (a *anobiiFileReader) Read(anobiiFile *os.File) ([]*Anobii, error) {
	var anobiiBooks []*Anobii
	if err := gocsv.UnmarshalFile(anobiiFile, &anobiiBooks); err != nil {
		return nil, err
	}

	return anobiiBooks, nil
}
