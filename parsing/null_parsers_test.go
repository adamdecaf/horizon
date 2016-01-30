package parsing

import (
	"testing"
)

func TestNullSingleParser(t *testing.T) {
	parser := NullSingleParser{}
	res, err := parser.Parse("hello")
	if err != nil {
		t.Fatalf("found error with null single parser err=%s", err)
	}
	if res != "" {
		t.Fatal("expected an empty string as null single parse result got=%s", res)
	}
}

func TestNullMultiParser(t *testing.T) {
	parser := NullMultiParser{}
	res, err := parser.Parse("hi there")
	if err != nil {
		t.Fatalf("found error with null multi parser err=%s", err)
	}
	if len(res) > 0 {
		t.Fatal("expected an empty string as null multi parse result got=%s", res)
	}
}
