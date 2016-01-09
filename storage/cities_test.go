package storage

import (
	"testing"
	"github.com/adamdecaf/horizon/utils"
)

func TestReadWriteCity(t *testing.T) {
	name := utils.RandString(20)
	empty, err := SearchCitiesByName(name)

	if err != nil {
		t.Fatalf("error reading city when we expected to see no results = %s", err)
	}

	if len(empty) != 0 {
		t.Fatal("found cities some how when we didn't expect to find any")
	}

	id := utils.RandString(20)
	stateId := utils.RandString(36)
	city := City{id, name, stateId}

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
