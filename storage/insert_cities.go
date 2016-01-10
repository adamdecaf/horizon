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

func InsertRawCitiesFromStates() (int64, error) {
	fmt.Println("insert cities")

	states, err := ReadAllStates()
	if err != nil {
		return 0, err
	}

	var count int64

	file, err := os.Open("./storage/raw-data/cities")
	if err != nil {
		return 0, err
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
			id := utils.UUID()
			city_name := utils.StripQuotesAndTrim(row[CITY_NAME])
			state_name := utils.StripQuotesAndTrim(row[STATE_INDEX])

			existing, err := SearchCitiesByName(city_name)
			if err != nil {
				fmt.Printf("[Storage/insert] error reading city %s\n", city_name)
				return 0, err
			} else {
				if len(existing) == 0 {
					// states
					for i := range states {
						if state_name == states[i].Name {
							// only insert city if we don't find one (and we find the state)
							city := City{id, city_name, states[i].Id}

							written := WriteCity(city)
							if written != nil {
								fmt.Printf("[Storage] error inserting raw city %s, %s, (err=%s)\n", id, city_name, *written)
								return 0, err
							} else {
								count = count + 1
							}
						}
					}
				}
			}
		}
	}

	return count, nil
}
