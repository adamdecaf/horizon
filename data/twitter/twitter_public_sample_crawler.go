package twitter

import (
	"net/url"
	"log"
	"time"
	"github.com/ChimeraCoder/anaconda"
	"github.com/adamdecaf/horizon/utils"
	"github.com/adamdecaf/horizon/data"
	"github.com/adamdecaf/horizon/metrics"
)

var tweets_meter = metrics.Meter("twitter.tweets")
var insert_failed_meter = metrics.Meter("twitter.insert-failed")

type TwitterPublicSampleCrawler struct {
	data.Crawler
}

func (c TwitterPublicSampleCrawler) Run() *error {
	log.Println("TwitterPublicSampleCrawler")

	api, err := create_twitter_api()
	if err != nil {
                log.Printf("error when creating twitter api err=%s\n", err)
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
			twitter_user := TwitterUser{}

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
			basic_tweet := BasicTweet{}

			parsed_tweet_time, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)

			if err != nil {
				log.Printf("Error parsing tweet date time (value='%s') (err=%s)\n", tweet.CreatedAt, err)
				continue
			}

			basic_tweet.CreatedAt = parsed_tweet_time
			basic_tweet.Id = tweet.IdStr
			basic_tweet.Text = tweet.Text
			basic_tweet.User = twitter_user

			go store_tweet_and_user(basic_tweet, twitter_user)

			go SpawnTwitterParsers(basic_tweet)
		}
	}

	return nil
}

func store_tweet_and_user(basic_tweet BasicTweet, twitter_user TwitterUser) {
	if err := WriteTwitterTweet(basic_tweet); err != nil {
		log.Printf("error while writing twitter tweet err=%s\n", *err)
		insert_failed_meter.Mark(1)
	} else {
		tweets_meter.Mark(1)

		// ignore duplicates, so we don't check for failures
		WriteTwitterUser(twitter_user)
	}
}

func SpawnTwitterPublicSampleCrawler() *error {
	config := utils.NewConfig()

	if run := config.Get("TWITTER_PUBLIC_CRAWLER_ENABLED"); run == "yes" {
		log.Println("[Spawn] creating TwitterPublicSampleCrawler")
		crawler := TwitterPublicSampleCrawler{}
		return data.RunCrawler(crawler)
	}
	return nil
}
