package retrieval

import (
	"log"
)

type Crawler interface {
	Run() *error
}

func RunCrawler(crawler Crawler) *error {
	if err := crawler.Run(); err != nil {
		log.Printf("error in crawler run err=%s\n", *err)
		return err
	}
	return nil
}
