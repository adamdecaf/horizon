package retrieval

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
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

func SpawnRedditCrawler() (Crawler, error) {
	return RedditCrawler{}, nil
}
