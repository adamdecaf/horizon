package storage

import (
	"fmt"
	"time"
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

func SearchTwitterTweetsById(tweet_id string) (BasicTweet, error) {
	var id string
	var text string
	var twitter_user_id string
	var created_at time.Time

	db, err := InitializeStorage()
	if err != nil {
		return BasicTweet{}, err
	}

	defer db.Close()

	rows, err := db.Query("select tweet_id, twitter_user_id, text, created_at from twitter_tweets where tweet_id=$1 limit 1;", tweet_id)

	if err != nil {
		return BasicTweet{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &twitter_user_id, &text, &created_at)
		if err != nil {
			fmt.Printf("[Storage] error getting tweet = %s\n", err)
			return BasicTweet{}, err
		}

		user, err := SearchTwitterUserById(twitter_user_id)
		if err != nil {
			return BasicTweet{}, err
		}

		return BasicTweet{created_at, id, text, user}, nil
	}

	return BasicTweet{}, fmt.Errorf("Unable to find tweet for id %s", tweet_id)
}

func SearchTwitterUserById(twitter_user_id string) (TwitterUser, error) {
	var id string
	var name string
	var screen_name string
	var created_at time.Time

	db, err := InitializeStorage()
	if err != nil {
		return TwitterUser{}, err
	}

	defer db.Close()

	rows, err := db.Query("select twitter_user_id, name, screen_name, created_at from twitter_users where twitter_user_id=$1 limit 1;", twitter_user_id)

	if err != nil {
		return TwitterUser{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name, &screen_name, &created_at)
		if err != nil {
			fmt.Printf("[Storage] error getting twitter user = %s\n", err)
			return TwitterUser{}, err
		}

		return TwitterUser{created_at, id, name, screen_name}, nil
	}

	return TwitterUser{}, fmt.Errorf("Unable to find twitter user for id %s", twitter_user_id)
}

func WriteTwitterTweet(tweet BasicTweet) *error {
	db, err := InitializeStorage()
	if err != nil {
		return &err
	}

	defer db.Close()

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
	db, err := InitializeStorage()
	if err != nil {
		return &err
	}

	defer db.Close()

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
