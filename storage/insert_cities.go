package storage

import (
	"bufio"
	"encoding/csv"
	"io"
	"fmt"
	"os"
	"github.com/adamdecaf/horizon/utils"
)

const (
	CITY_NAME = 0
	STATE_INDEX = 1
)

func InsertRawCitiesFromStates() *error {
	fmt.Println("insert cities")

	states, err := ReadAllStates()
	if err != nil {
		return &err
	}

	file, err := os.Open("./storage/raw-data/cities")
	if err != nil {
		return &err
	}

	reader := csv.NewReader(bufio.NewReader(file))

	defer file.Close()

	for {
		row, err := reader.Read()

		// stop on EOF
		if err == io.EOF {
			break
		}

		if row[CITY_NAME] != "" {
			city_name := utils.StripQuotesAndTrim(row[CITY_NAME])
			state_name := utils.StripQuotesAndTrim(row[STATE_INDEX])

			for i := range states {
				if state_name == states[i].Name {
					go write_city(city_name, states[i].Id)
				}
			}
		}
	}

	return nil
}

func write_city(city_name string, state_id string) {
	existing, err := SearchCitiesByNameAndState(city_name, state_id)
	if err != nil {
		fmt.Printf("[Storage/insert] error reading city %s (err=%s)\n", city_name, err)
		return
	}

	if len(existing) == 0 {
		// only insert city if we don't find one (and we find the state)
		id := utils.UUID()
		city := City{id, city_name, state_id}
		written := WriteCity(city)
		if written != nil {
			fmt.Printf("[Storage] error inserting raw city %s, %s, (err=%s)\n", id, city_name, *written)
		}
	}
}
