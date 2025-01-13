package goodreads

import "fmt"

type GoodReads struct {
	Title                   string `csv:"Title"`
	Author                  string `csv:"Author"`
	ISBN                    string `csv:"ISBN"`
	MyRating                string `csv:"My Rating"`
	AverageRating           string `csv:"Average Rating"`
	Publisher               string `csv:"Publisher"`
	Binding                 string `csv:"Binding"`
	YearPublished           string `csv:"Year Published"`
	OriginalPublicationYear string `csv:"Original Publication Year"`
	DateRead                string `csv:"Date Read"`
	DateAddedShelves        string `csv:"Date Added,Shelves"`
	Shelves                 string `csv:"Shelves"`
	Bookshelves             string `csv:"Bookshelves"`
	MyReview                string `csv:"My Review"`
}

func (a GoodReads) String() string {
	return fmt.Sprintf("ISBN: %s, Title: %s, Author: %s",
		a.ISBN,
		a.Title,
		a.Author)
}
