package twitter

import (
	"log"
	"os"
	"github.com/adamdecaf/horizon/data"
)

// how is this translated to ChimeraCoder/anaconda ?
// https://dev.twitter.com/rest/reference/get/users/lookup

// Scan through twitter users we have and find:
// - # of followers (recorded with timestamp)
// - # some public profile details
// Save off ^^ into table(s)
// - twitter_user_followers
// - twitter_user_details

// self-rate limited and
// as large of batches as possible
// add metrics

type TwitterUserCrawler struct {
	data.Crawler
}

func (c TwitterUserCrawler) Run() *error {
	_, err := create_twitter_api()
	if err != nil {
		return &err
	}

	// maybe?
	// api.EnableThrottling(5 * time.Second, 20)

	// grab twitter user ids from table.
	// which ones haven't we grabbed yet? (table with api requests we make)

	// params := url.Values{}
	// params.Add("language", "en") // add user ids

	// Make the call (in goroutine?)
	// api.GetUsersLookupByIds(ids []int64, v url.Values) (u []User, err error)

	// handle parsing of json in goroutine (so we don't block?)
	// but maybe block http call to sync up?

	// or have goroutines with mutex waiting to grab from postgres

	return nil
}

func SpawnTwitterUserCrawler() *error {
	if run := os.Getenv("TWITTER_USER_CRAWLER_ENABLED"); run == "yes" {
		log.Println("[Spawn] creating TwitterUserCrawler")
		crawler := TwitterUserCrawler{}
		return data.RunCrawler(crawler)
	}
	return nil
}
