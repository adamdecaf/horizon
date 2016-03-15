package twitter

import (
	"log"
	"github.com/adamdecaf/horizon/metrics"
	parsers "github.com/adamdecaf/horizon/data/parsers"
	// "github.com/zhenjl/porter2" // use to stem tweets
)

var url_parser = parsers.UrlMultiParser{}

var tweets_with_urls_meter = metrics.Meter("twitter.tweets-with-urls")

func SpawnTwitterParsers(tweet BasicTweet) *error {
	// find urls
	err := find_and_store_urls(tweet)
	if err != nil {
		return err
	}
	return nil
}

func find_and_store_urls(tweet BasicTweet) *error {
	urls, err := url_parser.Parse(tweet.Text)
	if err != nil {
		log.Printf("Error when parsing tweet err=%s\n", err)
		return &err
	}

	if len(urls) > 0 {
		tweets_with_urls_meter.Mark(1)
		err := WriteTwitterUrls(tweet.Id, urls)
		if err != nil {
			return err
		}
	}

	return nil
}
