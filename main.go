package main

import (
	"log"
	"sync"
	"github.com/adamdecaf/horizon/data"
	postgres "github.com/adamdecaf/horizon/data/engines/postgres"
	internet "github.com/adamdecaf/horizon/data/internet"
	reddit "github.com/adamdecaf/horizon/data/reddit"
	twitter "github.com/adamdecaf/horizon/data/twitter"
	wordcount "github.com/adamdecaf/horizon/data/twitter/word-count"
	"github.com/adamdecaf/horizon/metrics"
)

func main() {
	log.Println("starting horizon")

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
	go twitter.SpawnTwitterMentionProcessor()
	go wordcount.SpawnWordCountReprocessor()

	// wait forever
	block()
}

func block() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Println("waiting for exit signal")
	wg.Wait()
}
