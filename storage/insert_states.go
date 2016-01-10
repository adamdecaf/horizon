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
	STATE_NAME = 0
	STATE_ABBREVIATION = 1
)

func InsertRawStates() (int64, error) {
	var count int64

	file, err := os.Open("./storage/raw-data/states")
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

		if row[STATE_NAME] != "" {
			id := utils.UUID()
			name := utils.StripQuotesAndTrim(row[STATE_NAME])
			abbr := utils.StripQuotesAndTrim(row[STATE_ABBREVIATION])

			existing, err := SearchStatesByName(name)
			if err != nil {
				fmt.Printf("[Storage/insert] error reading state %s\n", name)
				return 0, err
			} else {
				if len(existing) == 0 {
					// only insert state if we don't fine one already
					state := State{id, name, abbr}
					written := WriteState(state)
					if written != nil {
						fmt.Printf("[Storage] error inserting raw state %s, %s, %s, (err=%s)\n", id, name, abbr, *written)
						return 0, err
					} else {
						count = count + 1
					}
				}
			}
		}
	}

	return count, nil
}
