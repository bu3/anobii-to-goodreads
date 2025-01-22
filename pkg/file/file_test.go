package file_test

import (
	"github.com/bu3/anobii-to-goodreads/pkg/file"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	It("should read a file", func() {
		fileManager := file.New()

		readFile, err := fileManager.ReadFile("../../testdata/anobii-exported.csv")
		Expect(err).ToNot(HaveOccurred())

		Expect(readFile).To(Not(BeNil()))
		Expect(len(readFile)).Should(BeNumerically(">", 0))
	})

	It("should return an error if file does not exist", func() {
		fileManager := file.New()

		readFile, err := fileManager.ReadFile("../../testdata/unknown-file")
		Expect(err).To(HaveOccurred())

		Expect(readFile).To(BeNil())
	})
})
