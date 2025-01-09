package acceptance_test

import (
	"bufio"
	"fmt"
	"github.com/bu3/anobii-to-goodreads/convert"
	"github.com/bu3/anobii-to-goodreads/mapping"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"os"
	"strings"
)

const generatedFilePath = "../testdata/goodreads_generated_output.csv"

var _ = Describe("Convert", Ordered, func() {
	Context("Anobii to Goodreads", func() {
		It("should match expected file", func() {
			mapper := mapping.New()
			converter := convert.Converter{
				Mapper: mapper,
			}

			expectedContent, err := readAndNormalize("../testdata/goodreads_expected.csv")

			err = converter.Convert("../testdata/anobii-exported.csv", generatedFilePath)
			Expect(err).ToNot(HaveOccurred())

			generatedFileContent, err := readAndNormalize(generatedFilePath)
			if err != nil {
				log.Fatal(err)
			}

			Expect(string(expectedContent)).To(Equal(string(generatedFileContent)))
		})

		AfterAll(func() {
			deleteGeneratedFile()
		})
	})
})

func deleteGeneratedFile() {
	if _, err := os.Stat(generatedFilePath); err == nil {
		// File exists, proceed to delete
		err = os.Remove(generatedFilePath)
		if err != nil {
			fmt.Printf("Error deleting file: %v\n", err)
			return
		}
		fmt.Printf("File %s deleted successfully.\n", generatedFilePath)
	} else if os.IsNotExist(err) {
		// File does not exist
		fmt.Printf("File %s does not exist.\n", generatedFilePath)
	} else {
		// An unexpected error occurred
		fmt.Printf("Error checking file: %v\n", err)
	}
}

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
