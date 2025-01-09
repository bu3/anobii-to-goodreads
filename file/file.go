package file

import "os"

type Manager interface {
	ReadFile(path string) ([]byte, error)
}

type fileManager struct{}

func (f fileManager) ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func New() Manager {
	return &fileManager{}
}
