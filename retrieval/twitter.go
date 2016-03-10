package retrieval

import (
	"fmt"
	"os"
	"github.com/ChimeraCoder/anaconda"
)

func create_twitter_api() (*anaconda.TwitterApi, error) {
	consumer_key := os.Getenv("TWITTER_CONSUMER_KEY")
	consumer_secret_key := os.Getenv("TWITTER_CONSUMER_SECRET")

	if consumer_key == "" || consumer_secret_key == "" {
		err := fmt.Errorf("[Retrieval] Missing consumer keys (key=%s) (secret=%s)", consumer_key, consumer_secret_key)
		return nil, err
	}

	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_secret := os.Getenv("TWITTER_ACCESS_SECRET")

	if access_token == "" || access_secret == "" {
		err := fmt.Errorf("[Retrieval] Missing access tokens (token=%s) (secret=%s)", access_token, access_secret)
		return nil, err
	}

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret_key)

	api := anaconda.NewTwitterApi(access_token, access_secret)
	return api, nil
}
