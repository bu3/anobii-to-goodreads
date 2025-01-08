package goodreads

import (
	"github.com/gocarina/gocsv"
	"os"
)

func Write(outputFile string, items *[]GoodReads) error {
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
