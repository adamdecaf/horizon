package retrieval

import (
	"fmt"
)

type Crawler interface {
	Run() *error
}

func RunCrawler(crawler Crawler) *error {
	if err := crawler.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
