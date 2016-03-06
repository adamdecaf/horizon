package parsing

import (
	"log"
	"github.com/adamdecaf/horizon/metrics"
	"github.com/adamdecaf/horizon/storage"
)

var url_parser = UrlMultiParser{}

var tweets_with_urls_meter = metrics.Meter("twitter.tweets-with-urls")

func SpawnTwitterParsers(tweet storage.BasicTweet) *error {
	// find urls
	urls, err := url_parser.Parse(tweet.Text)
	if err != nil {
		log.Printf("Error when parsing tweet err=%s\n", err)
		return &err
	}

	if len(urls) > 0 {
		tweets_with_urls_meter.Mark(1)
		err := storage.WriteTwitterUrls(tweet.Id, urls)
		if err != nil {
			return err
		}
	}

	return nil
}
