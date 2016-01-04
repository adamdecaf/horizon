package retrieval

import (
	"fmt"
)

type NullCrawler struct {
	Crawler
}

func (c NullCrawler) Run() *error {
	fmt.Println("NullCrawler.Run()")
	return nil
}

func SpawnNullCrawler() *error {
	fmt.Println("[Spawn] NullCrawler")
	crawler := NullCrawler{}
	return RunCrawler(crawler)
}
