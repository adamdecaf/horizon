package twitter

import (
	"fmt"
	"log"
	"time"
	postgres "github.com/adamdecaf/horizon/data/engines/postgres"
	"github.com/adamdecaf/horizon/utils"
)

type TwitterUser struct {
	CreatedAt time.Time
	Id string
	Name string
	ScreenName string
}

type BasicTweet struct {
	CreatedAt time.Time
	Id string
	Text string
	User TwitterUser
}

type TwitterMentionProcessorRun struct {
	Id string
	RangeStart time.Time
	RangeEnd time.Time
	CreatedAt time.Time
}

func SearchTwitterUserById(twitter_user_id string) (TwitterUser, error) {
	var id string
	var name string
	var screen_name string
	var created_at time.Time

	db, err := postgres.InitializePostgres()
	if err != nil {
		return TwitterUser{}, err
	}

	rows, err := db.Query("select twitter_user_id, name, screen_name, created_at from twitter_users where twitter_user_id=$1 limit 1;", twitter_user_id)

	if err != nil {
		return TwitterUser{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name, &screen_name, &created_at)
		if err != nil {
			log.Printf("[Storage] error getting twitter user = %s\n", err)
			return TwitterUser{}, err
		}

		return TwitterUser{created_at, id, name, screen_name}, nil
	}

	return TwitterUser{}, fmt.Errorf("Unable to find twitter user for id %s", twitter_user_id)
}

func SearchTwitterTweetsById(tweet_id string) (BasicTweet, error) {
	base := "select tweet_id, twitter_user_id, text, created_at from twitter_tweets where tweet_id=$1 limit 1;"
	res, err := query_twitter_tweets(base, tweet_id)

	if err != nil {
		return BasicTweet{}, err
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return BasicTweet{}, fmt.Errorf("Unable to find tweet tweet_id=%s", tweet_id)
}

func GrabTweetsViaDateRange(start time.Time, end time.Time, max int) ([]BasicTweet, error) {
	base := "select tweet_id, twitter_user_id, text, created_at from twitter_tweets where created_at >= $1 and created_at <= $2 limit $3;"
	return query_twitter_tweets(base, start, end, max)
}

func query_twitter_tweets(base string, rest ...interface{}) ([]BasicTweet, error) {
	var id string
	var text string
	var twitter_user_id string
	var created_at time.Time

	var results []BasicTweet

	db, err := postgres.InitializePostgres()

	if err != nil {
		return results, err
	}

	rows, err := db.Query(base, rest...)

	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &twitter_user_id, &text, &created_at)
		if err != nil {
			log.Printf("[Storage] error getting tweet = %s\n", err)
			return results, err
		}

		user, err := SearchTwitterUserById(twitter_user_id)
		if err != nil {
			return results, err
		}

		tweet := BasicTweet{created_at, id, text, user}
		results = append(results, tweet)
	}

	return results, nil
}

func WriteTwitterTweet(tweet BasicTweet) *error {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into twitter_tweets (tweet_id, twitter_user_id, text, created_at) values ($1, $2, $3, $4)", tweet.Id, tweet.User.Id, tweet.Text, tweet.CreatedAt)
	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert twitter tweet as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}

func WriteTwitterUser(user TwitterUser) *error {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into twitter_users (twitter_user_id, name, screen_name, created_at) values ($1, $2, $3, $4)", user.Id, user.Name, user.ScreenName, user.CreatedAt)
	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert twitter user as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}

func WriteTwitterMentionProcessorRun(run TwitterMentionProcessorRun) *error {
	db, err := postgres.InitializePostgres()

	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into twitter_mention_processing_runs (twitter_mention_processing_id, range_start, range_end, created_at) values ($1, $2, $3, $4)", run.Id, run.RangeStart, run.RangeEnd, run.CreatedAt)
	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert twitter mention processor result as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}

func WriteTwitterUrls(tweet_id string, urls []string) *error {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	for i := range urls {
		id := utils.UUID()
		result, err := db.Exec("insert into twitter_tweet_urls (twitter_tweet_url_id, tweet_id, url) values ($1, $2, $3)", id, tweet_id, urls[i])
		if err != nil {
			return &err
		}

		rows, err := result.RowsAffected()
		if err != nil {
			return &err
		}

		if rows != 1 {
			err := fmt.Errorf("[Storage] didn't insert twitter urls as expected (rows=%s)\n", rows)
			return &err
		}
	}

	return nil
}
