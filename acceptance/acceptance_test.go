package acceptance_test

import (
	"github.com/bu3/anobii-to-goodreads/convert"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"os"
)

var _ = Describe("Convert", func() {
	Context("Anobii to Goodreads", func() {
		XIt("should match expected file", func() {
			expectedContent, err := os.ReadFile("../testdata/goodreads_expected.csv")
			if err != nil {
				log.Fatal(err)
			}

			generatedFile := "../testdata/goodreads_generated_output.csv"
			err = convert.Convert("../testdata/anobii-exported.csv", generatedFile)
			Expect(err).ToNot(HaveOccurred())

			generatedFileContent, err := os.ReadFile(generatedFile)
			if err != nil {
				log.Fatal(err)
			}

			Expect(string(expectedContent)).To(Equal(string(generatedFileContent)))
		})
	})

})
