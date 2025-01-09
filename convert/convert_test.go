package convert_test

import (
	"fmt"
	"github.com/bu3/anobii-to-goodreads/convert"
	"github.com/bu3/anobii-to-goodreads/providers/anobii"
	"github.com/bu3/anobii-to-goodreads/providers/goodreads"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"os"
)

var _ = Describe("Convert", func() {
	Context("Anobii to Goodreads", func() {
		It("should convert one to another", func() {
			//TODO Move this code somewhere
			file, err := os.CreateTemp("/tmp", "prefix")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(file.Name())

			fileManager := &mockFileManager{
				mockOpenFile: func(path string) (*os.File, error) {
					return file, nil
				},
			}

			mockAnobiiReader := &mockAnobiiReader{
				mockRead: func(anobiiFile *os.File) ([]*anobii.Anobii, error) {
					return []*anobii.Anobii{}, nil
				},
			}

			mockMapper := &MockMapper{
				mockMapItem: func(input *anobii.Anobii) (goodreads.GoodReads, error) {
					return goodreads.GoodReads{}, nil
				},
				mockMapList: func(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error) {
					goodReads := []goodreads.GoodReads{
						goodreads.GoodReads{},
						goodreads.GoodReads{},
					}
					return &goodReads, nil
				},
			}
			//TODO Move this code somewhere

			converter := convert.Converter{
				FileManager:  fileManager,
				Mapper:       mockMapper,
				AnobiiReader: mockAnobiiReader,
			}

			Expect(mockAnobiiReader.readInvocations).To(Equal(0))
			Expect(fileManager.openFileInvocations).To(Equal(0))
			Expect(mockMapper.mapItemInvocations).To(Equal(0))
			Expect(mockMapper.mapListInvocations).To(Equal(0))

			converter.Convert("path1", "path2")

			Expect(mockAnobiiReader.readInvocations).To(Equal(1))
			Expect(fileManager.openFileInvocations).To(Equal(1))
			Expect(mockMapper.mapListInvocations).To(Equal(1))
			Expect(mockMapper.mapItemInvocations).To(Equal(2))
		})
	})
})

//TODO Move this code somewhere

type mockAnobiiReader struct {
	readInvocations int
	mockRead        func(anobiiFile *os.File) ([]*anobii.Anobii, error)
}

func (r *mockAnobiiReader) Read(anobiiFile *os.File) ([]*anobii.Anobii, error) {
	r.readInvocations++
	return r.mockRead(anobiiFile)
}

type mockFileManager struct {
	openFileInvocations int
	mockOpenFile        func(path string) (*os.File, error)
}

func (f *mockFileManager) OpenFile(filePath string) (*os.File, error) {
	f.openFileInvocations++
	fmt.Println("opening file", filePath, f.openFileInvocations)
	return f.mockOpenFile(filePath)
}

func (f *mockFileManager) Reset() {
	f.openFileInvocations = 0
}

type MockMapper struct {
	mapListInvocations int
	mapItemInvocations int

	mockMapList func(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error)
	mockMapItem func(input *anobii.Anobii) (goodreads.GoodReads, error)
}

func (m *MockMapper) Reset() {
	m.mapListInvocations = 0
	m.mapItemInvocations = 0
}

func (m *MockMapper) MapItem(input *anobii.Anobii) (goodreads.GoodReads, error) {
	m.mapItemInvocations++
	return m.mockMapItem(input)
}

func (m *MockMapper) MapList(inputs []*anobii.Anobii) (*[]goodreads.GoodReads, error) {
	m.mapListInvocations++
	var outputs []goodreads.GoodReads
	for _, input := range inputs {
		item, _ := m.MapItem(input)
		outputs = append(outputs, item)
	}
	return &outputs, nil
}
