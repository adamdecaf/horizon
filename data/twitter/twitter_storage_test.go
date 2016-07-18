package twitter

import (
	"time"
	"testing"
	"github.com/adamdecaf/horizon/utils"
)

func TestReadWriteTwitterData(t *testing.T) {
	// Test Data
	tweet_id := utils.RandString(20)
	twitter_user_id := utils.RandString(20)
	name := utils.RandString(20)
	screen_name := utils.RandString(20)
	text := utils.RandString(100)
	created_at := time.Now()

	// Methods
	empty_tweet, err := SearchTwitterTweetsById(tweet_id)
	if err == nil {
		t.Fatalf("we expected to find an error when reading an empty tweet")
	}
	if empty_tweet.Id != "" {
		t.Fatalf("why did we find a tweet? (tweet_id = %s)", tweet_id)
	}

	empty_user, err := SearchTwitterUserById(twitter_user_id)
	if err == nil {
		t.Fatalf("we expected to find an error when reading an empty twitter user")
	}
	if empty_user.Id != "" {
		t.Fatalf("why did we find a user? (twitter_user_id = %s)", twitter_user_id)
	}

	// Data
	user := TwitterUser{created_at, twitter_user_id, name, screen_name}
	tweet := BasicTweet{created_at, tweet_id, text, user}

	// Insert
	if err := WriteTwitterTweet(tweet); err != nil {
		t.Fatalf("had an error when writing tweet=%s, err=%s", tweet, *err)
	}

	if err := WriteTwitterUser(user); err != nil {
		t.Fatalf("had an error when writing twitter user=%s, err=%s", user, *err)
	}

	// Check
	found_user, err := SearchTwitterUserById(twitter_user_id)
	if err != nil {
		t.Fatalf("found an error when searching for twitter users err=%s", err)
	}
	if found_user.Id != user.Id {
		t.Fatalf("the inserted user doesn't match what we tried to write")
	}

	found_tweet, err := SearchTwitterTweetsById(tweet_id)
	if err != nil {
		t.Fatalf("found an error when searching for tweet err=%s", err)
	}
	if found_tweet.Id != tweet.Id {
		t.Fatalf("found tweet and created tweet don't match (found_tweet=%s, tweet_id = %s)", found_tweet, tweet_id)
	}

	// Grab tweets from a minute ago.
	start := time.Now().Add(-1 * time.Minute)
	end := time.Now()
	results, err := GrabTweetsViaDateRange(start, end, 10)
	if err != nil {
		t.Fatalf("found error when getting date range of tweets err=%s", err)
	}
	if len(results) < 1 {
		t.Fatalf("found no tweets when we expected some err=%s", err)
	}
}

func TestWriteProcessorMentionRun(t *testing.T) {
	id := utils.RandString(20)
	now := time.Now()
	result := TwitterMentionProcessorRun{id, now, now, now}
	if err := WriteTwitterMentionProcessorRun(result); err != nil {
		t.Fatalf("erorr writing twitter mention processor run err=%s", *err)
	}
}
