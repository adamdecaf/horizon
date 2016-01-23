package search

import (
	"regexp"
	"strings"
)

func ParseSearchFlags(input string) map[string]string {
	normalized := strings.TrimSpace(input)

	if normalized == "" {
		return nil
	}

	pairs := regexp.MustCompile(`(?i)([\w]+:'[\s-!#-&\w]+')|([\w]+:[\w]+)`).FindAllString(input, 5)

	var terms map[string]string
	terms = make(map[string]string)

	for i := range pairs {
		splits := strings.Split(pairs[i], ":")
		if len(splits) > 1 {
			joined := strings.Join(splits[1:len(splits)], "")
			stripped := strings.TrimSuffix(strings.TrimPrefix(joined, "'"), "'")
			terms[splits[0]] = stripped
		}
	}

	return terms
}
