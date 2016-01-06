package storage

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func TestReadWriteCity(t *testing.T) {
	name := RandString(20)
	empty, err := SearchCitiesByName(name)

	if err != nil {
		t.Fatalf("error reading city when we expected to see no results = %s", err)
	}

	if len(empty) != 0 {
		t.Fatal("found cities some how when we didn't expect to find any")
	}

	id := RandString(20)
	city := City{id, name}

	if written := WriteCity(city); written != nil {
		t.Fatalf("error when writing city name=%s, err=%s", name, *written)
	}

	cities, err := SearchCitiesByName(name)

	if err != nil {
		t.Fatalf("error finding city that should exist name=%s, err=%s\n", name, err)
	}

	if len(cities) == 1 {
		found := cities[0]
		if found.Id != city.Id || found.Name != city.Name {
			t.Fatalf("cities don't match (written=%s) (found=%s)", city, found)
		}
	} else {
		t.Fatalf("found multiple cities when we expected one name=%s", name)
	}
}
