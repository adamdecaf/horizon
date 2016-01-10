package utils

import (
	"strings"
)

func StripQuotesAndTrim(str string) string {
	no_quotes := strings.Replace(str, "'", "", -1)
	return strings.TrimSpace(no_quotes)
}
