package twitter

import (
	"log"
	"time"
	"github.com/adamdecaf/horizon/data"
	"github.com/adamdecaf/horizon/utils"
)

type MentionReprocessRange struct {
	RangeStart time.Time
	RangeEnd time.Time
}

type TwitterMentionProcessor struct {
	data.ReProcessor
}

var mention_reprocess_step = 1 * time.Hour
var max_tweets_per_mention_reprocess = 5000

func (p TwitterMentionProcessor) Run() *error {
	log.Printf("starting TwitterMentionProcessor")
	reprocess_range, err := get_tweet_range()
	if err != nil {
		return &err
	}

	tweets, err := GrabTweetsViaDateRange(reprocess_range.RangeStart, reprocess_range.RangeEnd, max_tweets_per_mention_reprocess)
	if err != nil {
		return &err
	}

	for i := range tweets {
		// todo: parsing for mentions
		// todo: write found mentions to table
		log.Printf(tweets[i].Text)
	}

	run := TwitterMentionProcessorRun{}
	run.Id = utils.UUID()
	run.RangeStart = reprocess_range.RangeStart
	run.RangeEnd = reprocess_range.RangeEnd
	run.CreatedAt = time.Now()

	if ok := WriteTwitterMentionProcessorRun(run); ok != nil {
		return ok
	}

	return nil
}

func get_tweet_range() (MentionReprocessRange, error) {
	// for now, hand configured
	reprocess_range := MentionReprocessRange{}
	reprocess_range.RangeStart = time.Now().Add(7 * 24 * time.Hour)
	reprocess_range.RangeEnd = time.Now()

	return reprocess_range, nil
}

func SpawnTwitterMentionProcessor() *error {
	config := utils.NewConfig()

	if run := config.Get("TWITTER_MENTION_PROCESSOR_ENABLED"); run == "yes" {
		log.Printf("[Spawn] TwitterMentionProcessor (run=%s)\n", run)
		processor := TwitterMentionProcessor{}
		return data.StartReProcessor(processor)
	}

	return nil
}
