package storage

import (
	"strings"
	"github.com/adamdecaf/horizon/utils"
)

func InsertHostnames() (int64, error) {
	var inserted_hostnames int64

	blob, err := utils.GzipDecompressFile("./storage/raw-data/top-1m-hostnames.gz")
	if err != nil {
		return 0, err
	}

	split := strings.Split(blob, "\n")
	for i := range split {
		hostname := utils.StripQuotesAndTrim(split[i])

		if hostname != "" {
			_, err := SearchHostnameByValue(hostname)
			if err != nil {
				id := utils.UUID()
				if ok := WriteHostname(Hostname{id, hostname}); ok != nil {
					return 0, err
				}

				inserted_hostnames = inserted_hostnames + 1
			}
		}
	}

	return inserted_hostnames, nil
}
