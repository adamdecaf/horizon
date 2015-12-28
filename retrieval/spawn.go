package retrieval

import (
	"fmt"
)

func SpawnCrawlers() {
	fmt.Println("Spawning crawlers")

	// null crawler
	fmt.Println("[Spawn] NullCrawler")
	crawler, err := SpawnNullCrawler()
	if err != nil {
		panic(err)
	}
	CheckCrawlerRun(crawler)
}

func CheckCrawlerRun(crawler Crawler) {
	if err := crawler.Run(); err != nil {
		fmt.Println(err)
	}
}
