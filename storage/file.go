package storage

import (
	"fmt"
)

type FileStorage interface {
	Save([]byte) (string, error)
}

func InitFileStorage() (FileStorage, error) {
	return LocalFileStorage{}, nil
}

// local file storage
type LocalFileStorage struct {
	FileStorage
}

func (f LocalFileStorage) Save([]byte) (string, error) {
	fmt.Println("[Storage] save to LocalFileStorage")
	return "", nil
}
