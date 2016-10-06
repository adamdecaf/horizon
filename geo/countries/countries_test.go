package countries

import (
	"testing"
	"github.com/adamdecaf/horizon/utils"
)

func TestReadAndWriteCountries(t *testing.T) {
	name := utils.RandString(20)
	empty, err := SearchCountryByName(name)

	if err != nil {
		t.Fatalf("error reading country when we expected no results = %s", err)
	}

	if len(empty) != 0 {
		t.Fatalf("why did we find a country? (name=%s)", name)
	}

	id := utils.RandString(20)
	country := Country{id, name}

	if written := WriteCountry(country); written != nil {
		t.Fatalf("error when writing country name=%s, err=%s", name, *written)
	}

	countries, err := SearchCountryByName(name)

	if err != nil {
		t.Fatalf("error when finding country by name=%s, err=%s", name, err)
	}

	if len(countries) == 1 {
		found := countries[0]
		if found.Id != country.Id || found.Name != country.Name {
			t.Fatalf("countries don't match (written=%s) (found=%s)", country, found)
		}
	} else {
		t.Fatalf("found multiple countries when we expected one name=%s", name)
	}

	single_country, err := FindCountryById(country.Id)
	if err != nil {
		t.Fatalf("didn't find single country because of error (err=%s)\n", err)
	}

	if single_country == nil {
		t.Fatalf("unable to find single country... (country_id = %s)", country.Id)
	}
}
