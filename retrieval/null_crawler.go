package retrieval

import (
	"log"
)

type NullCrawler struct {
	Crawler
}

func (c NullCrawler) Run() *error {
	log.Println("NullCrawler.Run()")
	return nil
}

func SpawnNullCrawler() *error {
	log.Println("[Spawn] NullCrawler")
	crawler := NullCrawler{}
	return RunCrawler(crawler)
}
