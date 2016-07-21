package wordcount

import (
	"fmt"
	"log"
	"time"
	twitter "github.com/adamdecaf/horizon/data/twitter"
	postgres "github.com/adamdecaf/horizon/data/engines/postgres"
)

type HourlyWordCount struct {
	Word string // lowercase
	Count int
	Hour time.Time
}

func hasWordCountsForHour(hour time.Time) (bool, error) {
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

func getTweetChunk(hour time.Time, offset, limit int) []twitter.TextOnlyTweet {
	next := hour.Add(1 * time.Hour)
	tweets, err := twitter.GetTweetTextFromDateRange(hour, next, offset, limit)
	if err != nil {
		log.Printf("error reading tweets err=%s\n", err)
		return nil
	}
	return tweets
}

func storeTweetWordCounts(counts []HourlyWordCount) *error {
	if len(counts) == 0 {
		return nil
	}

	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	base := "insert into twitter_hourly_word_counts (word, count, hour) values ($1,$2,$3)"
	for i := range counts {
		c := counts[i]
		result, err := db.Exec(base, c.Word, c.Count, c.Hour)
		if err != nil {
			return &err
		}

		rows, err := result.RowsAffected()
		if err != nil {
			return &err
		}

		if rows != 1 {
			err := fmt.Errorf("[Storage] didn't write twitter word counts as expected (rows=%s)\n", rows)
			return &err
		}
	}

	return nil
}
