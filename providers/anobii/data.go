package anobii

import "fmt"

type Anobii struct {
	ISBN              string `csv:"ISBN"`
	Title             string `csv:"Title"`
	Subtitle          string `csv:"Subtitle"`
	Author            string `csv:"Author"`
	Format            string `csv:"Format"`
	NumberOfPages     string `csv:"Number of pages"`
	Publisher         string `csv:"Publisher"`
	DateOfPublication string `csv:"Date of publication"`
	PrivateNotes      string `csv:"Private notes"`
	CommentTitle      string `csv:"Comment Title"`
	CommentContent    string `csv:"Comment Content"`
	ReadingStatus     string `csv:"Reading status"`
	Vote              string `csv:"Vote"`
	Tags              string `csv:"Tags"`
}

func (a Anobii) String() string {
	return fmt.Sprintf("ISBN: %q, Title: %q, Subtitle: %q, Author: %q",
		a.ISBN,
		a.Title,
		a.Subtitle,
		a.Author)
}
