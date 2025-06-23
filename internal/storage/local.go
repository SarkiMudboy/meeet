package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ObjectStorage interface {
	Save(name, path string, file *os.File) (string, error)
}

type LocalStorage struct {
	root string
}

func NewLocalStorage(root string) *LocalStorage {

	return &LocalStorage{
		root: root,
	}
}

func (l *LocalStorage) Save(name, path string, file *os.File) (fullPath string, err error) {
	// save path: full path to storage location

	cwd, _ := os.Getwd()
	fullPath = filepath.Join(cwd, l.root, path, name)

	diskFile, err := os.Create(path)
	if err != nil {
		fmt.Printf("%s\n%s\n", fullPath, err)
		return
	}

	defer diskFile.Close()
	_, err = io.Copy(diskFile, file)
	if err != nil {
		return
	}

	return
}
