package storage

import (
	"log"
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
	log.Println("[Storage] save to LocalFileStorage")
	return "", nil
}
