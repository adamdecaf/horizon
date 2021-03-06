package states

import (
	"fmt"
	"log"
	"strings"
	"github.com/adamdecaf/horizon/data/engines/postgres"
)

type State struct {
	Id string `json:"stateId"`
	Name string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

func FindStateById(state_id string) (*State, error) {
	res, err := query_states("select state_id, name, abbreviation from states where state_id=$1 limit 1;", state_id)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		state := res[0]
		return &state, nil
	}
	return nil, fmt.Errorf("[Storage] unable to find state by state_id %s", state_id)
}

func SearchStatesByName(raw string) ([]State, error) {
	return query_states("select state_id, name, abbreviation from states where lower(name) like '%' || $1 || '%';", strings.ToLower(raw))
}

func ReadAllStates() ([]State, error) {
	return query_states("select state_id, name, abbreviation from states;")
}

func query_states(base string, rest ...interface{}) ([]State, error) {
	states := make([]State, 0)

	var id string
	var name string
	var abbreviation string

	db, err := postgres.InitializePostgres()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(base, rest...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &abbreviation)
		if err != nil {
			log.Printf("[Storage] error getting state = %s\n", err)
		}
		states = append(states, State{id, name, abbreviation})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return states, nil
}

func WriteState(state State) *error {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into states (state_id, name, abbreviation) values ($1, $2, $3)", state.Id, state.Name, state.Abbreviation)

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
