package main

import (
	"log"
	"github.com/adamdecaf/horizon/analysis"
	"github.com/adamdecaf/horizon/data"
	postgres "github.com/adamdecaf/horizon/data/engines/postgres"
	internet "github.com/adamdecaf/horizon/data/internet"
	reddit "github.com/adamdecaf/horizon/data/reddit"
	twitter "github.com/adamdecaf/horizon/data/twitter"
	"github.com/adamdecaf/horizon/metrics"
)

func main() {
	log.Println("Starting horizon")

	// Setup postgres tables
	postgres.MigrateStorage()

	// async things
	go data.InsertData()
	go metrics.InitializeStdoutReporter()

	// crawlers
	go reddit.SpawnRedditCrawler()
	go twitter.SpawnTwitterPublicSampleCrawler()
	go internet.SpawnWhoisCrawler()

	// reprocessors
	log.Println("[reprocessors] Spawning reprocessors")
	go twitter.SpawnTwitterMentionProcessor()

	// Start the analysis http server
	analysis.StartHttpServer()
}
