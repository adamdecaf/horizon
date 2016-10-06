package internet

import (
	"fmt"
	"log"
	"github.com/adamdecaf/horizon/data/engines/postgres"
)

type Hostname struct {
	Id string `json:"hostnameId"`
	Value string `json:"value"`
}

func SearchHostnameByValue(value string) (Hostname, error) {
	res, err := query_hostnames("select hostname_id, value from hostnames where value=$1 limit 1;", value)
	if err != nil {
		return Hostname{}, err
	}

	if len(res) > 0 {
		hostname := res[0]
		return hostname, nil
	}
	return Hostname{}, fmt.Errorf("[Storage] unable to find hostname by value %s", value)
}

func query_hostnames(base string, rest ...interface{}) ([]Hostname, error) {
	hostnames := make([]Hostname, 0)

	var id string
	var value string

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
		err := rows.Scan(&id, &value)
		if err != nil {
			log.Printf("[Storage] error getting hostname = %s\n", err)
		}
		hostnames = append(hostnames, Hostname{id, value})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hostnames, nil
}

func WriteHostname(hostname Hostname) *error {
	db, err := postgres.InitializePostgres()
	if err != nil {
		return &err
	}

	result, err := db.Exec("insert into hostnames (hostname_id, value) values ($1, $2)", hostname.Id, hostname.Value)

	if err != nil {
		return &err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return &err
	}

	if rows != 1 {
		err := fmt.Errorf("[Storage] didn't insert hostname as expected (rows=%s)\n", rows)
		return &err
	}

	return nil
}
