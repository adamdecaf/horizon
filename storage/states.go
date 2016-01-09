package storage

import (
	"fmt"
	"strings"
)

type State struct {
	Id string `json:"stateId"`
	Name string `json:"name"`
}

func SearchStatesByName(raw string) ([]State, error) {
	states := make([]State, 0)

	var id string
	var name string

	db, err := InitializeStorage()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select state_id, name from states where lower(name) like '%' || $1 || '%';", strings.ToLower(raw))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("[Storage] error getting state = %s\n", err)
		}
		states = append(states, State{id, name})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return states, nil
}

func WriteState(state State) *error {
	db, err := InitializeStorage()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into states (state_id, name) values ($1, $2)", state.Id, state.Name)
	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert state as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}
