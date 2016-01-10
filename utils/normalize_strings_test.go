package utils

import (
	"testing"
)

func TestQuotedStringRemoval(t *testing.T) {
	nothing := "aaa  "
	has_quotes := "'aaa'"
	has_quotes_and_whitespace := "  'aaa' "

	if res := StripQuotesAndTrim(nothing); res != "aaa" {
		t.Fatalf("%s did not equal expected %s\n", res, "aaa")
	}

	if res := StripQuotesAndTrim(has_quotes); res != "aaa" {
		t.Fatalf("%s did not equal expected %s\n", res, "aaa")
	}

	if res := StripQuotesAndTrim(has_quotes_and_whitespace); res != "aaa" {
		t.Fatalf("%s did not equal expected %s\n", res, "aaa")
	}
}
