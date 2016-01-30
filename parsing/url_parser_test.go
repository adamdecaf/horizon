package parsing

import (
	"reflect"
	"testing"
)

func TestUrlParserEmptyString(t *testing.T) {
	input := ""

	parser := UrlMultiParser{}
	res, err := parser.Parse(input)
	if err != nil {
		t.Fatalf("found error while parsing '%s', err=%s", input, err)
	}
	if len(res) != 0 {
		t.Fatalf("res should be empty!")
	}
}

func TestUrlParserWithoutUrls(t *testing.T) {
	input := ""

	parser := UrlMultiParser{}
	res, err := parser.Parse(input)
	if err != nil {
		t.Fatalf("found error while parsing '%s', err=%s", input, err)
	}
	if len(res) != 0 {
		t.Fatalf("res should be empty!")
	}
}

func TestUrlParserWithUrls(t *testing.T) {
	input := "hi there foo.com hey there"
	answer := []string{"foo.com"}

	parser := UrlMultiParser{}
	res, err := parser.Parse(input)
	if err != nil {
		t.Fatalf("found error while parsing '%s', err=%s", input, err)
	}
	if !reflect.DeepEqual(res, answer) {
		t.Fatalf("unable to match answer (%s) to result (%s)", answer, res)
	}
}
