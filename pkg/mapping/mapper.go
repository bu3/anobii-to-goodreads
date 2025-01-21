package mapping

import (
	"github.com/bu3/anobii-to-goodreads/pkg/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/pkg/providers/goodreads"
	"regexp"
	"strings"
)

type AnobiiToGoodReadsMapper struct{}

func (m *AnobiiToGoodReadsMapper) MapItem(input *anobii.Anobii) (goodreads.GoodReads, error) {
	return goodreads.GoodReads{
		Title:    input.Title,
		ISBN:     input.ISBN,
		Author:   input.Author,
		MyRating: input.Vote,
		DateRead: parseDateRead(input.ReadingStatus),
		Shelves:  parseStatus(input.ReadingStatus),
	}, nil
}

func (m *AnobiiToGoodReadsMapper) MapList(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error) {
	var outputs []goodreads.GoodReads
	for _, input := range inputs {
		item, _ := m.MapItem(input)
		outputs = append(outputs, item)
	}
	return &outputs, nil
}

func parseStatus(readingStatus string) string {
	status := "to-read"
	if strings.Contains(strings.ToLower(readingStatus), "finished") {
		status = "read"
	}

	if strings.Contains(strings.ToLower(readingStatus), "being read") {
		status = "reading"
	}

	if strings.Contains(strings.ToLower(readingStatus), "abandoned") {
		status = "abandoned"
	}

	return status
}

func parseDateRead(readingStatus string) string {
	longDate := parseDateWithRegex(readingStatus, `[a-z A-Z]+\d{4}-\d{2}-\d{2}$`, `\d{4}-\d{2}-\d{2}$`)
	if len(longDate) > 0 {
		return longDate
	}
	shortDate := parseDateWithRegex(readingStatus, `[a-z A-Z]+\d{4}$`, `\d{4}$`)
	if len(shortDate) > 0 {
		return shortDate + "-01-01"
	}
	return ""
}

func parseDateWithRegex(value string, longRegex string, shortRegex string) string {
	statusWithTenCharacterLongDate := regexp.MustCompile(longRegex)
	if len(statusWithTenCharacterLongDate.FindStringSubmatch(value)) > 0 {
		dateRegex := regexp.MustCompile(shortRegex)
		matches := dateRegex.FindStringSubmatch(value)
		if len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}
