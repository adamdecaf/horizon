package utils

import (
	"testing"
)

func TestUUID(t *testing.T) {
	res := UUID()
	if len(res) != 36 {
		t.Fatalf("generated uuid (%s) is not 36 characters", res)
	}
}
