package mapping

import (
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
)

type AnobiiToGoodReadsMapper interface {
	MapItem(input *anobii.Anobii) (goodreads.GoodReads, error)
	MapList(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error)
}

func New() AnobiiToGoodReadsMapper {
	return &anobiiToGoodReadsMapper{}
}

type anobiiToGoodReadsMapper struct{}

func (m *anobiiToGoodReadsMapper) MapItem(input *anobii.Anobii) (goodreads.GoodReads, error) {
	return goodreads.GoodReads{
		Title:  input.Title,
		ISBN:   input.ISBN,
		Author: input.Author,
	}, nil
}

func (m *anobiiToGoodReadsMapper) MapList(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error) {
	var outputs []goodreads.GoodReads
	for _, input := range inputs {
		item, _ := m.MapItem(input)
		outputs = append(outputs, item)
	}
	return &outputs, nil
}
