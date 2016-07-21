package wordcount

import (
	"log"
	"math/rand"
	"strings"
	"time"
	configs "github.com/adamdecaf/horizon/configs"
	"github.com/adamdecaf/horizon/data"
	"github.com/adamdecaf/horizon/metrics"
	postgres "github.com/adamdecaf/horizon/data/engines/postgres"
)

var (
	tweets_processed = metrics.Meter("twitter.word-count.tweets-processed")
	words_counted = metrics.Meter("twitter.word-count.words-counted")
)

type WordCountReprocessor struct {
	data.ReProcessor
}

func (w WordCountReprocessor) Run() *error {
       for { 
	hour := randomHourStart()
	hasCounts, err := hasWordCountsForHour(hour)
	if err != nil {
		return &err
	}
        log.Printf("hasCounts = %b\n", hasCounts)
	if !hasCounts {
		offset := 0
		limit := 1000

		counts := make([]HourlyWordCount, 0)

		// generate counts
		for {
			tweets := getTweetChunk(hour, offset, limit)
                        log.Printf("tweets = %d\n", tweets)
			if len(tweets) == 0 {
				return nil
			}

			offset += len(tweets)

			for i := range tweets {
				tweets_processed.Mark(1)
				tweet := tweets[i]
				words := strings.Split(tweet.Text, " ")
				words_counted.Mark(int64(len(words)))

				subCounts := make(map[string]int, 0)
				for i := range words {
					subCounts[strings.ToLower(words[i])] += 0
				}

				for k,v := range subCounts {
					c := HourlyWordCount{k, v, hour}
					counts = append(counts, c)
				}
			}
		}

		// store counts
		return storeTweetWordCounts(counts)
	}
}
	return nil
}

func randomHourStart() time.Time {
	currMonth := int64(time.Now().Hour())
	month := time.Month(rand.Int63n(currMonth))
	day := int(rand.Int63n(30))
	if month == time.February {
		day = int(rand.Int63n(28))
	}
	hour := int(rand.Int63n(24))

	t := time.Date(2016, month, day, hour, 0, 0, 0, time.UTC)
	return t
}

func HasWordCountsForHour(hour time.Time) (bool, error) {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return false, err
	}

	rows, err := db.Query("select hour from twitter_hourly_word_counts where hour=$1 limit 1;", hour)
	defer rows.Close()

	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func SpawnWordCountReprocessor() *error {
	config := configs.NewConfig()
	if run := config.Get("TWITTER_WORD_COUNT_REPROCESSOR_ENABLED"); run == "yes" {
		log.Printf("[Spawn] Twitter WordCountReprocessor (run=%s)\n", run)
		reprocessor := WordCountReprocessor{}
		return data.StartReProcessor(reprocessor)
	}
	return nil
}
