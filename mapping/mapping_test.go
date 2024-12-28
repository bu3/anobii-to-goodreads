package mapping_test

import (
	"github.com/bu3/anobii-to-goodreads/mapping"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mapping", func() {
	Context("Anobii to Goodreads", func() {
		It("should map", func() {
			anobiiItem := anobii.Anobii{
				Title: "Foo Bar",
			}
			goodReads, err := mapping.AnobiiToGoodReads(anobiiItem)
			Expect(err).ToNot(HaveOccurred())
			Expect(goodReads).ToNot(BeNil())
			Expect(goodReads.Title).To(Equal(anobiiItem.Title))
		})
	})

})
