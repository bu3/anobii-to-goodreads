package file

import "os"

type Manager interface {
	OpenFile(path string) (*os.File, error)
}

type fileManager struct{}

func (f fileManager) OpenFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return file, nil
}

func New() Manager {
	return &fileManager{}
}
