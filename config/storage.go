package config

import (
	"io"
	"os"
)

type ObjectStorage interface {
	Save(file *os.File) (string, error)
}

type LocalStorage struct {
	root string
}

func NewLocalStorage(root string) *LocalStorage {

	return &LocalStorage{
		root: root,
	}
}

func (l *LocalStorage) Save(fileName string, file *os.File) (path string, err error) {

	diskFile, err := os.Create(l.root + fileName)
	if err != nil {
		return
	}

	defer diskFile.Close()
	_, err = io.Copy(diskFile, file)
	if err != nil {
		return
	}

	return
}
