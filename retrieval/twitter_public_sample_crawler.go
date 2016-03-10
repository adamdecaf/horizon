package retrieval

import (
	"net/url"
	"log"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/adamdecaf/horizon/metrics"
	"github.com/adamdecaf/horizon/storage"
	"github.com/adamdecaf/horizon/parsing"
)

var tweets_meter = metrics.Meter("twitter.tweets")
var insert_failed_meter = metrics.Meter("twitter.insert-failed")

type TwitterPublicSampleCrawler struct {
	Crawler
}

func (c TwitterPublicSampleCrawler) Run() *error {
	log.Println("TwitterPublicSampleCrawler")

	api, err := create_twitter_api()
	if err != nil {
		return &err
	}

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
				log.Printf("Error parsing user date time (value='%s') (err=%s)\n", tweet.User.CreatedAt, err)
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
				log.Printf("Error parsing tweet date time (value='%s') (err=%s)\n", tweet.CreatedAt, err)
				continue
			}

			basic_tweet.CreatedAt = parsed_tweet_time
			basic_tweet.Id = tweet.IdStr
			basic_tweet.Text = tweet.Text
			basic_tweet.User = twitter_user

			if err := storage.WriteTwitterTweet(basic_tweet); err != nil {
				log.Printf("error while writing twitter tweet err=%s\n", *err)
				insert_failed_meter.Mark(1)
			} else {
				tweets_meter.Mark(1)

				// ignore duplicates, so we don't check for failures
				storage.WriteTwitterUser(twitter_user)
			}

			go parsing.SpawnTwitterParsers(basic_tweet)
		}
	}

	return nil
}

func SpawnTwitterPublicSampleCrawler() *error {
	if run := os.Getenv("TWITTER_PUBLIC_CRAWLER_ENABLED"); run == "yes" {
		log.Println("[Spawn] creating TwitterPublicSampleCrawler")
		crawler := TwitterPublicSampleCrawler{}
		return RunCrawler(crawler)
	}
	return nil
}
