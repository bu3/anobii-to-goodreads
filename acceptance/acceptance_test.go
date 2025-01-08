package acceptance_test

import (
	"bufio"
	"github.com/bu3/anobii-to-goodreads/convert"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"os"
	"strings"
)

var _ = Describe("Convert", func() {
	Context("Anobii to Goodreads", func() {
		It("should match expected file", func() {
			expectedContent, err := readAndNormalize("../testdata/goodreads_expected.csv")
			if err != nil {
				log.Fatal(err)
			}

			generatedFile := "../testdata/goodreads_generated_output.csv"
			err = convert.Convert("../testdata/anobii-exported.csv", generatedFile)
			Expect(err).ToNot(HaveOccurred())

			generatedFileContent, err := readAndNormalize(generatedFile)
			if err != nil {
				log.Fatal(err)
			}

			Expect(string(expectedContent)).To(Equal(string(generatedFileContent)))
		})
	})

})

func readAndNormalize(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n") // Normalize line endings to '\n'
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return []byte(builder.String()), nil
}
