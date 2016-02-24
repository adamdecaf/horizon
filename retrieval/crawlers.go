package retrieval

import (
	"fmt"
)

func SpawnCrawlers() {
	fmt.Println("[retrieval] Spawning crawlers")

	go SpawnNullCrawler()
	go SpawnRedditCrawler()
	go SpawnTwitterPublicSampleCrawler()
	go SpawnWhoisCrawler()
}
