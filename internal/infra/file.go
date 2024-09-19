package infra

import (
	"io"
	"os"
	"path/filepath"
)

type (
	FileHandler struct {
		root string
	}
)

func NewFileHandler(root string) *FileHandler {
	return &FileHandler{
		root: root,
	}
}

func (h FileHandler) Create(filename string) (io.Writer, error) {
	path := h.buildCompletePath(filename)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	return os.Create(path)
}

func (h FileHandler) Exist(filename string) (bool, error) {
	path := h.buildCompletePath(filename)
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func (h FileHandler) buildCompletePath(filename string) string {
	return filepath.Join(h.root, filename)
}

func DefaultRootPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "Tabs"), nil
}
