package reprocess

import (
	"log"
	"os"
)

type TwitterMentionProcessor struct {
	Processor
}

func (p TwitterMentionProcessor) Run() *error {
	log.Printf("")



	return nil
}

func SpawnTwitterMentionProcessor() *error {
	if run := os.Getenv("TWITTER_MENTION_PROCESSOR_ENABLED"); run == "yes" {
		log.Printf("[Spawn] TwitterMentionProcessor (run=%s)\n", run)
		processor := TwitterMentionProcessor{}
		return StartProcessor(processor)
	}

	return nil
}
