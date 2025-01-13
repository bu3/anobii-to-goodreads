package goodreads

import (
	"github.com/bu3/anobii-to-goodreads/file"
	"github.com/gocarina/gocsv"
	"os"
)

type GoodreadsFile interface {
	Write(outputFile string, items *[]GoodReads) error
}

type goodReadsFile struct {
	fileManager file.Manager
}

func New(manager file.Manager) GoodreadsFile {
	return &goodReadsFile{
		fileManager: manager,
	}
}

func (g goodReadsFile) Write(outputFile string, items *[]GoodReads) error {
	goodReadsFile, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer goodReadsFile.Close()

	err = gocsv.MarshalFile(items, goodReadsFile)
	if err != nil {
		return err
	}
	return nil
}
