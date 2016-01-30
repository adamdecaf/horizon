package retrieval

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/adamdecaf/horizon/storage"
	"github.com/adamdecaf/horizon/parsing"
)

type TwitterPublicSampleCrawler struct {
	Crawler
}

func (c TwitterPublicSampleCrawler) Run() *error {
	fmt.Println("TwitterPublicSampleCrawler")

	consumer_key := os.Getenv("TWITTER_CONSUMER_KEY")
	consumer_secret_key := os.Getenv("TWITTER_CONSUMER_SECRET")

	if consumer_key == "" || consumer_secret_key == "" {
		err := fmt.Errorf("[Retrieval] Missing consumer keys (key=%s) (secret=%s)", consumer_key, consumer_secret_key)
		return &err
	}

	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_secret := os.Getenv("TWITTER_ACCESS_SECRET")

	if access_token == "" || access_secret == "" {
		err := fmt.Errorf("[Retrieval] Missing access tokens (token=%s) (secret=%s)", access_token, access_secret)
		return &err
	}

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret_key)

	api := anaconda.NewTwitterApi(access_token, access_secret)
	api.EnableThrottling(5 * time.Second, 20)

	params := url.Values{}
	params.Add("language", "en")

	res := api.PublicStreamSample(params)

	for {
		item := <-res.C
		tweet, ok := item.(anaconda.Tweet)
		if ok {
			// twitter user
			twitter_user := storage.TwitterUser{}

			parsed_user_time, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.User.CreatedAt)

			if err != nil {
				fmt.Printf("Error parsing user date time (value='%s') (err=%s)\n", tweet.User.CreatedAt, err)
				continue
			}

			twitter_user.CreatedAt = parsed_user_time
			twitter_user.Id = tweet.User.IdStr
			twitter_user.Name = tweet.User.Name
			twitter_user.ScreenName = tweet.User.ScreenName


			// tweet
			basic_tweet := storage.BasicTweet{}

			parsed_tweet_time, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)

			if err != nil {
				fmt.Printf("Error parsing tweet date time (value='%s') (err=%s)\n", tweet.CreatedAt, err)
				continue
			}

			basic_tweet.CreatedAt = parsed_tweet_time
			basic_tweet.Id = tweet.IdStr
			basic_tweet.Text = tweet.Text
			basic_tweet.User = twitter_user

			if err := storage.WriteTwitterTweet(basic_tweet); err != nil {
				fmt.Printf("error while writing twitter tweet err=%s\n", *err)

				// ignore duplicates, so we don't check for failures
				storage.WriteTwitterUser(twitter_user)
			}

			go parsing.SpawnTwitterParsers(basic_tweet)
		}
	}

	return nil
}

func SpawnTwitterPublicSampleCrawler() *error {
	if run := os.Getenv("TWITTER_CRAWLER_ENABLED"); run == "yes" {
		fmt.Println("[Spawn] creating TwitterPublicSampleCrawler")
		crawler := TwitterPublicSampleCrawler{}
		return RunCrawler(crawler)
	}

	return nil
}
