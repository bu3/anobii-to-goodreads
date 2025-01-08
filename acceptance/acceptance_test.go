package acceptance_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"os"
)

var _ = Describe("Convert", func() {
	Context("Anobii to Goodreads", func() {
		It("should match expected file", func() {
			expectedContent, err := os.ReadFile("../testdata/goodreads_expected.csv")
			if err != nil {
				log.Fatal(err)
			}

			Expect(string(expectedContent)).To(Equal("bar"))
		})
	})

})
