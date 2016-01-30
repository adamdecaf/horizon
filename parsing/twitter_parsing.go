package parsing

import (
	"fmt"

	"github.com/adamdecaf/horizon/storage"
)

var url_parser = UrlMultiParser{}

func SpawnTwitterParsers(tweet storage.BasicTweet) *error {
	// find urls
	urls, err := url_parser.Parse(tweet.Text)
	if err != nil {
		fmt.Printf("Error when parsing tweet err=%s\n", err)
		return &err
	}

	if len(urls) > 0 {
		err := storage.WriteTwitterUrls(tweet.Id, urls)
		if err != nil {
			return err
		}
	}

	return nil
}
