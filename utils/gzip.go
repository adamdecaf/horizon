package utils

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

func GzipDecompressFile(path string) (string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return "", err
	}

	gzip_reader, err := gzip.NewReader(reader)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(gzip_reader)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
