package convert

import (
	"github.com/bu3/anobii-to-goodreads/file"
	"github.com/bu3/anobii-to-goodreads/mapping"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
)

type Converter struct {
	FileManager  file.Manager
	Mapper       mapping.AnobiiToGoodReadsMapper
	AnobiiReader anobii.AnobiiFileReader
}

func (c Converter) Convert(inputFile string, outputFile string) error {
	inboundFile, err := c.FileManager.OpenFile(inputFile)
	if err != nil {
		panic(err)
	}
	defer inboundFile.Close()

	anobiiBooks, err := c.AnobiiReader.Read(inboundFile)
	if err != nil {
		panic(err)
	}

	c.Mapper.MapList(anobiiBooks)
	//if err != nil {
	//	panic(err)
	//}
	////goodreads.Write(outputFile, goodReadsItems)

	return nil
}
