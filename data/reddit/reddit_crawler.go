package reddit

import (
	"log"
	"github.com/adamdecaf/horizon/configs"
	"github.com/adamdecaf/horizon/data"
	"github.com/jzelinskie/geddit"
)

type RedditCrawler struct {
	data.Crawler
}

func (c RedditCrawler) Run() *error {
	log.Println("Starting RedditCrawler")

	config := configs.NewConfig()

	session, err := geddit.NewLoginSession(
		config.Get("REDDIT_USERNAME"),
		config.Get("REDDIT_PASSWORD"),
		"horizon retrieval agent",
	)

	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return &err
	}

	for _, s := range submissions {
		log.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
	}

	return nil
}

func SpawnRedditCrawler() *error {
	config := configs.NewConfig()

	if run := config.Get("REDDIT_CRAWLER_ENABLED"); run == "yes" {
		log.Printf("[Spawn] RedditCrawler (run=%s)\n", run)
		crawler := RedditCrawler{}
		return data.RunCrawler(crawler)
	}

	// todo
	// storage, err := storage.InitFileStorage()
	// if err != nil {
	// 	return &err
	// }

	// key, err := storage.Save(nil)
	// if err != nil {
	// 	return &err
	// }

	// log.Println("[Spawn] Reddit Crawler - key=" + key)

	return nil
}