package cities

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/adamdecaf/horizon/utils"
	"github.com/ivpusic/grpool"
)

const (
	CITY_NAME = 0
	STATE_INDEX = 1
)

func InsertRawCitiesFromStates(pool grpool.Pool) *error {
	states, err := ReadAllStates()
	if err != nil {
		return &err
	}

	file, err := os.Open("./data/raw-data/cities")
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
					go write_city(pool, city_name, states[i].Id)
				}
			}
		}
	}

	return nil
}

func write_city(pool grpool.Pool, city_name string, state_id string) {
	pool.JobQueue <- func() {
		defer pool.JobDone()

		existing, err := SearchCitiesByNameAndState(city_name, state_id)
		if err != nil {
			log.Printf("[Storage/insert] error reading city %s (err=%s)\n", city_name, err)
		}

		if len(existing) == 0 {
			// only insert city if we don't find one (and we find the state)
			id := utils.UUID()
			city := City{id, city_name, state_id}
			written := WriteCity(city)
			if written != nil {
				log.Printf("[Storage] error inserting raw city %s, %s, (err=%s)\n", id, city_name, *written)
			}
		}
	}
}
