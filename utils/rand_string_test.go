package utils

import (
	"testing"
)

func TestRandString(t *testing.T) {
	res := RandString(10)
	if len(res) != 10 {
		t.Fatalf("generated string (%s) is not 10 digits long", res)
	}
}
