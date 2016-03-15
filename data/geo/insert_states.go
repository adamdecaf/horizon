package geo

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
	STATE_NAME = 0
	STATE_ABBREVIATION = 1
)

func InsertRawStates(pool grpool.Pool) *error {
	file, err := os.Open("./storage/raw-data/states")
	if err != nil {
		return & err
	}

	reader := csv.NewReader(bufio.NewReader(file))

	defer file.Close()

	for {
		row, err := reader.Read()

		// stop on EOF
		if err == io.EOF {
			break
		}

		if row[STATE_NAME] != "" {
			name := utils.StripQuotesAndTrim(row[STATE_NAME])
			abbr := utils.StripQuotesAndTrim(row[STATE_ABBREVIATION])
			go write_state(pool, name, abbr)
		}
	}

	return nil
}

func write_state(pool grpool.Pool, name string, abbr string) {
	pool.JobQueue <- func() {
		defer pool.JobDone()

		existing, err := SearchStatesByName(name)
		if err != nil {
			log.Printf("[Storage/insert] error reading state %s\n", name)
		}
		if len(existing) == 0 {
			// only insert state if we don't fine one already
			id := utils.UUID()
			state := State{id, name, abbr}
			written := WriteState(state)
			if written != nil {
				log.Printf("[Storage] error inserting raw state %s, %s, %s, (err=%s)\n", id, name, abbr, *written)
			}
		}
	}
}
