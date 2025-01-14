package mapping_test

import (
	"github.com/bu3/anobii-to-goodreads/pkg/mapping"
	"github.com/bu3/anobii-to-goodreads/pkg/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/pkg/providers/goodreads"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mapping", func() {

	mapper := mapping.AnobiiToGoodReadsMapper{}

	Context("Anobii to Goodreads", func() {
		It("should map an item", func() {
			anobiiItem := anobii.Anobii{
				Title:         "Foo Bar",
				ISBN:          "1234",
				Author:        "John Doe",
				Vote:          "3",
				ReadingStatus: "Finished on 2023-03-30",
			}
			goodReadsItem, err := mapper.MapItem(&anobiiItem)
			Expect(err).ToNot(HaveOccurred())
			Expect(goodReadsItem).ToNot(BeNil())
			SingleItemExpectations(&anobiiItem, goodReadsItem)
		})

		It("should map a list of items", func() {
			mapper := mapping.AnobiiToGoodReadsMapper{}
			anobiiItem := anobii.Anobii{
				Title:         "Foo Bar",
				ISBN:          "1234",
				Author:        "John Doe",
				Vote:          "3",
				ReadingStatus: "Finished on 2023-03-30",
			}
			items := []*anobii.Anobii{&anobiiItem}
			goodReadsItems, err := mapper.MapList(items)

			Expect(err).ToNot(HaveOccurred())
			Expect(goodReadsItems).ToNot(BeNil())
			Expect(len(*goodReadsItems)).To(Equal(len(items)))

			for idx, goodReads := range *goodReadsItems {
				SingleItemExpectations(items[idx], goodReads)
			}
		})

		It("should map reading status", func() {
			data := map[string]string{
				"":                       "",
				"something odd":          "",
				"Finished on 2023-03-30": "2023-03-30",
				"Finished on 2023":       "2023-01-01",
			}
			for key, value := range data {
				anobiiItem := anobii.Anobii{
					ReadingStatus: key,
				}
				goodReadsItem, err := mapper.MapItem(&anobiiItem)
				Expect(err).ToNot(HaveOccurred())
				Expect(goodReadsItem).ToNot(BeNil())
				Expect(goodReadsItem.DateRead).To(Equal(value))
			}
		})
	})

})

func SingleItemExpectations(anobiiItem *anobii.Anobii, goodReads goodreads.GoodReads) {
	Expect(goodReads.Title).To(Not(BeEmpty()))
	Expect(goodReads.Title).To(Equal(anobiiItem.Title))
	Expect(goodReads.ISBN).To(Not(BeEmpty()))
	Expect(goodReads.ISBN).To(Equal(anobiiItem.ISBN))
	Expect(goodReads.Author).To(Not(BeEmpty()))
	Expect(goodReads.Author).To(Equal(anobiiItem.Author))
	Expect(goodReads.MyRating).To(Not(BeEmpty()))
	Expect(goodReads.MyRating).To(Equal(anobiiItem.Vote))
	Expect(goodReads.DateRead).To(Not(BeEmpty()))
	Expect(goodReads.DateRead).To(Equal("2023-03-30"))
}
