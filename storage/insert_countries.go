package storage

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"github.com/adamdecaf/horizon/utils"
)

func InsertCountries() (int64, error) {
	var count int64

	file, err := os.Open("./storage/raw-data/countries")
	if err != nil {
		return 0, err
	}

	reader := bufio.NewReader(file)

	defer file.Close()

	for {
		row, err := reader.ReadString('\n')

		// stop on EOF
		if err == io.EOF {
			break
		}

		name := utils.StripQuotesAndTrim(row)
		existing, err := SearchCountryByName(name)
		if err != nil {
			return 0, err
		}

		if len(existing) == 0 {
			id := utils.UUID()
			written := WriteCountry(Country{id, name})
			if written != nil {
				fmt.Printf("[Storage] error inserting country id=%s, name=%s, err=%s\n", id, name, *written)
				return 0, err
			}
			count = count + 1
		}
	}

	return count, nil
}
