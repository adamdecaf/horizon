package parsers

import (
	"github.com/mvdan/xurls"
)

type UrlMultiParser struct {
	MultiParser
}

func (p UrlMultiParser) Parse(input string) ([]string, error) {
	parsed := xurls.Relaxed.FindAllString(input, -1)
	return parsed, nil
}
