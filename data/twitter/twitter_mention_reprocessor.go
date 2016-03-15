package twitter

import (
	"log"
	"os"
	"github.com/adamdecaf/horizon/data"
)

type TwitterMentionReProcessor struct {
	data.ReProcessor
}

func (p TwitterMentionReProcessor) Run() *error {
	log.Printf("")

	return nil
}

func SpawnTwitterMentionReProcessor() *error {
	if run := os.Getenv("TWITTER_MENTION_PROCESSOR_ENABLED"); run == "yes" {
		log.Printf("[Spawn] TwitterMentionReProcessor (run=%s)\n", run)
		reprocessor := TwitterMentionReProcessor{}
		return data.StartReProcessor(reprocessor)
	}

	return nil
}
