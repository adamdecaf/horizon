package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// The default location for .horizon file
const DefaultHorizonConfigFilePath = ".horizon"

type Config struct {
	horizon_file_kv map[string]string
}

func (c Config) Get(name string) string {
	// Is it in .horizon file?
	if v := c.horizon_file_kv[name]; v != "" {
		return v
	}

	return os.Getenv(name)
}

func NewConfig() Config {
	return NewConfigFromPath(DefaultHorizonConfigFilePath)
}

func NewConfigFromPath(path string) Config {
	kv := parse_horizon_file(path)
	return Config{kv}
}

func parse_horizon_file(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		// file probably doesn't exist
		return nil
	}

	var kv map[string]string
	kv = make(map[string]string)

	reader := bufio.NewReader(file)
	defer file.Close()

	for {
		row, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		parts := strings.Split(row, "=")
		if len(parts) < 1 {
			log.Println("error parsing .horizon line -- not printing sensitive row from file")
			continue
		}

		key := parts[0]
		value := strings.TrimSpace(strings.Join(parts[1:], "="))

		kv[key] = value
	}

	return kv
}
