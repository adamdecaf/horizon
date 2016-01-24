package retrieval

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
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

	res := api.PublicStreamSample(url.Values{})
	// api.SetDelay(5 * time.Second)

	for i := 1; i <= 3; i++ {
		item := <-res.C
		tweet, ok := item.(anaconda.Tweet)
		if !ok {
			fmt.Printf("bad resp gotten from twitter (%s)\n", tweet.Text)
		} else {
			// twitter user
			twitter_user := TwitterUser{}

			parsed_user_time, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.User.CreatedAt)

			if err != nil {
				fmt.Printf("Error parsing user date time (value='%s') (err=%s)", tweet.User.CreatedAt, err)
				continue
			}

			twitter_user.CreatedAt = parsed_user_time
			twitter_user.Id = tweet.User.IdStr
			twitter_user.Name = tweet.User.Name
			twitter_user.ScreenName = tweet.User.ScreenName

			fmt.Println(twitter_user)

			// tweet
			basic_tweet := BasicTweet{}

			parsed_tweet_time, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)

			if err != nil {
				fmt.Printf("Error parsing tweet date time (value='%s') (err=%s)", tweet.CreatedAt, err)
				continue
			}

			basic_tweet.CreatedAt = parsed_tweet_time
			basic_tweet.Id = tweet.IdStr
			basic_tweet.Text = tweet.Text
			basic_tweet.User = twitter_user

			fmt.Println(basic_tweet)
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
