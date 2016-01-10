package main

import (
	"fmt"

	"github.com/adamdecaf/horizon/analysis"
	"github.com/adamdecaf/horizon/retrieval"
	"github.com/adamdecaf/horizon/storage"
)

func main() {
	fmt.Println("Starting horizon")

	// Setup tables
	storage.MigrateStorage()

	// Insert base data
	go storage.InsertData()

	// spawn crawlers
	go retrieval.SpawnNullCrawler()
	go retrieval.SpawnRedditCrawler()

	// Start the analysis http server
	analysis.StartHttpServer()
}
