package mapping_test

import (
	"github.com/bu3/anobii-to-goodreads/pkg/mapping"
	"github.com/bu3/anobii-to-goodreads/pkg/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/pkg/providers/goodreads"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mapping", func() {
	Context("Anobii to Goodreads", func() {
		It("should map an item", func() {
			mapper := mapping.AnobiiToGoodReadsMapper{}
			anobiiItem := anobii.Anobii{
				Title:  "Foo Bar",
				ISBN:   "1234",
				Author: "John Doe",
			}
			goodReadsItem, err := mapper.MapItem(&anobiiItem)
			Expect(err).ToNot(HaveOccurred())
			Expect(goodReadsItem).ToNot(BeNil())
			SingleItemExpectations(&anobiiItem, goodReadsItem)
		})

		It("should map a list of items", func() {
			mapper := mapping.AnobiiToGoodReadsMapper{}
			anobiiItem := anobii.Anobii{
				Title:  "Foo Bar",
				ISBN:   "1234",
				Author: "John Doe",
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
	})

})

func SingleItemExpectations(anobiiItem *anobii.Anobii, goodReads goodreads.GoodReads) {
	Expect(goodReads.Title).To(Equal(anobiiItem.Title))
	Expect(goodReads.ISBN).To(Equal(anobiiItem.ISBN))
	Expect(goodReads.Author).To(Equal(anobiiItem.Author))
}
