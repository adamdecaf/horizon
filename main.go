package main

import (
	"fmt"

	"github.com/adamdecaf/horizon/retrieval"
	"github.com/adamdecaf/horizon/storage"
)

func main() {
	fmt.Println("Starting horizon")

	// Setup tables
	storage.MigrateStorage()

	// Insert base data
	storage.InsertData()

	// spawn crawlers
	retrieval.SpawnRedditCrawler()
	retrieval.SpawnNullCrawler()
}
