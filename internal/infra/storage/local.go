package storage

import (
	"os"
	"path/filepath"
)

type LocalStorageAdapter struct {
	BasePath string
}

func (l *LocalStorageAdapter) Upload(path string, data []byte) error {
	fullPath := filepath.Join(l.BasePath, path)
	os.MkdirAll(filepath.Dir(fullPath), 0755)
	return os.WriteFile(fullPath, data, 0644)
}

func (l *LocalStorageAdapter) Download(path string) ([]byte, error) {
	fullPath := filepath.Join(l.BasePath, path)
	return os.ReadFile(fullPath)
}

func (l *LocalStorageAdapter) Delete(path string) error {
	fullPath := filepath.Join(l.BasePath, path)
	return os.Remove(fullPath)
}
