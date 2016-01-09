package storage

import (
	"fmt"
	"strings"
)

type City struct {
	Id string `json:"cityId"`
	Name string `json:"name"`
	StateId string `json:"stateId"`
}

func SearchCitiesByName(raw string) ([]City, error) {
	cities := make([]City, 0)

	var id string
	var name string
	var stateId string

	db, err := InitializeStorage()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select city_id, name, state_id from cities where lower(name) like '%' || $1 || '%';", strings.ToLower(raw))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &stateId)
		if err != nil {
			fmt.Printf("[Storage] error getting cities = %s\n", err)
		}
		cities = append(cities, City{id, name, stateId})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func WriteCity(city City) *error {
	db, err := InitializeStorage()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into cities (city_id, name, state_id) values ($1, $2, $3)", city.Id, city.Name, city.StateId)
	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert city as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}
