package storage

import (
	"fmt"
	"strings"

	"github.com/adamdecaf/horizon/utils"
	"github.com/ivpusic/grpool"
)

func InsertHostnames(pool grpool.Pool) *error {
	blob, err := utils.GzipDecompressFile("./storage/raw-data/top-1m-hostnames.gz")
	if err != nil {
		return &err
	}

	split := strings.Split(blob, "\n")
	for i := range split {
		hostname := utils.StripQuotesAndTrim(split[i])
		if hostname != "" {
			go write_hostname(pool, hostname)
		}
	}

	return nil
}

func write_hostname(pool grpool.Pool, hostname string) {
	pool.JobQueue <- func() {
		defer pool.JobDone()

		_, err := SearchHostnameByValue(hostname)
		if err != nil {
			id := utils.UUID()
			if err := WriteHostname(Hostname{id, hostname}); err != nil {
				fmt.Printf("[storage] error writing hostname '%s', err=%s\n", hostname, *err)
			}
		}
	}
}
