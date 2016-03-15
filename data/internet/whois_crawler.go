package internet

import (
	"fmt"
	"log"
	"github.com/adamdecaf/horizon/data"
	whois "github.com/adamdecaf/go-whois"
	"github.com/adamdecaf/go-whois/parsing"
)

type WhoisCrawler struct {
	data.Crawler
}

func (c WhoisCrawler) Run() *error {
	log.Println("WhoisCrawler run")

	// Grab some ready-to-run hostnames to lookup whois results for

	res, err := whois.WhoisQuery("google.com")
	if err != nil {
		return &err
	}

	fmt.Println(parsing.ParseWhoisResponse(res))

	return nil
}

func SpawnWhoisCrawler() *error {
	log.Println("[Spawn] WhoisCrawler")
	crawler := WhoisCrawler{}
	return data.RunCrawler(crawler)
}
