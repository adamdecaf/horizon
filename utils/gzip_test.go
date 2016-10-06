package utils

import (
	"testing"
)

func TestGzipReadEmptyFile(t *testing.T) {
	body, err := GzipDecompressFile("testdata/empty-file.gz")
	if err != nil {
		t.Fatalf("error reading compressed file err=%s", err)
	}
	if body != "" {
		t.Fatalf("Found something in 'empty-file.gz' == '%s'", body)
	}
}

func TestGzipReadFile(t *testing.T) {
	body, err := GzipDecompressFile("testdata/file-with-data.gz")
	if err != nil {
		t.Fatalf("error reading compressed file err=%s", err)
	}
	if body != "hello\n" {
		t.Fatalf("file-with-data contents don't match '%s' expected 'hello'", body)
	}
}
