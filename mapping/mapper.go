package mapping

import (
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
)

type AnobiiToGoodReadsMapper struct{}

func (m *AnobiiToGoodReadsMapper) MapItem(input anobii.Anobii) (goodreads.GoodReads, error) {
	return goodreads.GoodReads{
		Title: input.Title,
	}, nil
}

func (m *AnobiiToGoodReadsMapper) MapList(inputs []anobii.Anobii) ([]goodreads.GoodReads, error) {
	var outputs []goodreads.GoodReads
	for _, input := range inputs {
		item, _ := m.MapItem(input)
		outputs = append(outputs, item)
	}
	return outputs, nil
}
