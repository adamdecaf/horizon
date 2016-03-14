package main

import (
	"log"

	"github.com/adamdecaf/horizon/analysis"
	"github.com/adamdecaf/horizon/metrics"
	"github.com/adamdecaf/horizon/reprocess"
	"github.com/adamdecaf/horizon/retrieval"
	"github.com/adamdecaf/horizon/storage"
)

func main() {
	log.Println("Starting horizon")

	// Setup tables
	storage.MigrateStorage()

	// Insert base data
	go storage.InsertData()

	// spawn crawlers
	go retrieval.SpawnCrawlers()

	// start (re)processors
	go reprocess.SpawnProcessors()

	// Start stdout reporting
	go metrics.InitializeStdoutReporter()

	// Start the analysis http server
	analysis.StartHttpServer()
}
