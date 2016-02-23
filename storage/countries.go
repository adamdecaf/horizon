package storage

import (
	"fmt"
	"strings"
)

type Country struct {
	Id string `json:"countryId"`
	Name string `json:"name"`
}

func FindCountryById(country_id string) (*Country, error) {
	res, err := QueryCountriesTable("select country_id, name from countries where country_id = $1", country_id)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		country := res[0]
		return &country, nil
	}
	return nil, fmt.Errorf("[Storage] unable to find country by country_id %s", country_id)
}

func SearchCountryByName(raw string) ([]Country, error) {
	return QueryCountriesTable("select country_id, name from countries where lower(name) like '%' || $1 || '%';", strings.ToLower(raw))
}

func QueryCountriesTable(base string, rest ...interface{}) ([]Country, error) {
	countries := make([]Country, 0)

	var id string
	var name string

	db, err := InitializeStorage()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(base, rest...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("[Storage] error getting country = %s\n", err)
		}
		countries = append(countries, Country{id, name})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return countries, nil
}

func WriteCountry(country Country) *error {
	db, err := InitializeStorage()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into countries (country_id, name) values ($1, $2)", country.Id, country.Name)

	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert country as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}
