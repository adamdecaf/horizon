package retrieval

import (
	"log"
)

func SpawnCrawlers() {
	log.Println("[retrieval] Spawning crawlers")

	go SpawnNullCrawler()
	go SpawnRedditCrawler()
	go SpawnTwitterPublicSampleCrawler()
	go SpawnWhoisCrawler()
}
