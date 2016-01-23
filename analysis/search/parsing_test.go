package search

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestEmptySearchFlagParsing(t *testing.T) {
	res := ParseSearchFlags("")

	if res != nil {
		t.Fatalf("expected empty result of search flags from empty input")
	}
}

func TestValidSearchFlagParsing(t *testing.T) {
	res1 := ParseSearchFlags("city:ames name:adam")

	res1_ans := map[string]string{"city":"ames", "name":"adam"}
	if eq := reflect.DeepEqual(res1, res1_ans); !eq {
		t.Fatalf("res1 maps don't equal -- found %s", res1)
	}

	res2 := ParseSearchFlags("city:'des moines' name:adam")

	res2_ans := map[string]string{"city":"des moines", "name":"adam"}
	if eq := reflect.DeepEqual(res2, res2_ans); !eq {
		t.Fatalf("res2 maps don't equal -- found %s", res2)
	}
}

func TestInvalidSearchFlagParsing(t *testing.T) {
	res := ParseSearchFlags("city:")

	if len(res) != 0 {
		t.Fatalf("we expected an empty parse response")
	}
}

// sorta fuzz-testing
// from https://github.com/google/gofuzz/blob/master/fuzz.go
type charRange struct {
	first, last rune
}

func (r *charRange) choose(rand *rand.Rand) rune {
	count := int64(r.last - r.first)
	return r.first + rune(rand.Int63n(count))
}

var justAlphanumeric = []charRange{
	{'a', 'z'},
	{'A', 'Z'},
	{'0', '9'},
}

var largerCharRanges = []charRange{
	{' ', '!'},
	{'#', '&'},
	{'a', 'z'},
	{'A', 'Z'},
	{'0', '9'},
}

func randString(r *rand.Rand, ranges []charRange) string {
	n := r.Intn(30) + 1
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = ranges[r.Intn(len(ranges))].choose(r)
	}
	return string(runes)
}

func TestFuzzSearchFlagParsing(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {
		str1 := randString(rng, justAlphanumeric)
		str2 := randString(rng, largerCharRanges)

		query := "city:" + str1 + " name:'" + str2 + "'"

		res1 := ParseSearchFlags(query)

		res1_ans := map[string]string{"city":str1, "name":str2}
		if eq := reflect.DeepEqual(res1, res1_ans); !eq {
			t.Fatalf("res1 maps don't equal -- found %s -- query %s", res1, query)
		}
	}
}
