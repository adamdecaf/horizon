package retrieval

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
	"github.com/adamdecaf/horizon/storage"
)

type RedditCrawler struct {
	Crawler
}

func (c RedditCrawler) Run() *error {
	fmt.Println("Starting RedditCrawler")

	session, err := geddit.NewLoginSession(
		os.Getenv("REDDIT_USERNAME"),
		os.Getenv("REDDIT_PASSWORD"),
		"horizon retrieval agent",
	)

	if err != nil {
		fmt.Println(err)
		return &err
	}

	options := geddit.ListingOptions{
		Limit: 10,
	}

	// default frontpage
	submissions, err := session.DefaultFrontpage(geddit.DefaultPopularity, options)

	// personal frontpage
	// submissions, err = session.Frontpage(geddit.DefaultPopularity, options)

	if err != nil {
		fmt.Println(err)
		return &err
	}

	for _, s := range submissions {
		fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
	}

	return nil
}

func SpawnRedditCrawler() *error {
	if run := os.Getenv("REDDIT_CRAWLER_ENABLED"); run == "yes" {
		fmt.Printf("[Spawn] RedditCrawler (run=%s)\n", run)
		crawler := RedditCrawler{}
		return RunCrawler(crawler)
	}

	storage, err := storage.InitFileStorage()
	if err != nil {
		return &err
	}

	key, err := storage.Save(nil)
	if err != nil {
		return &err
	}

	fmt.Println("[Spawn] Reddit Crawler - key=" + key)

	return nil
}
