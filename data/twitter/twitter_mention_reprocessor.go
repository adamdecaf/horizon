package twitter

import (
	"log"
	"github.com/adamdecaf/horizon/configs"
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
	config := configs.NewConfig()

	if run := config.Get("TWITTER_MENTION_PROCESSOR_ENABLED"); run == "yes" {
		log.Printf("[Spawn] TwitterMentionReProcessor (run=%s)\n", run)
		reprocessor := TwitterMentionReProcessor{}
		return data.StartReProcessor(reprocessor)
	}

	return nil
}
