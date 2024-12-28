package mapping

import (
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
)

func AnobiiToGoodReads(input anobii.Anobii) (goodreads.GoodReads, error) {
	return goodreads.GoodReads{
		Title: input.Title,
	}, nil
}
