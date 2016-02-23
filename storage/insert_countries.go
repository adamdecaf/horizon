package storage

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"github.com/adamdecaf/horizon/utils"
)

func InsertCountries() *error {
	file, err := os.Open("./storage/raw-data/countries")
	if err != nil {
		return &err
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
		go write_country(name)
	}

	return nil
}


func write_country(name string) {
	existing, err := SearchCountryByName(name)
	if err != nil {
		fmt.Printf("[storage] Error searching for country by name '%s'", name)
	}

	if len(existing) == 0 {
		id := utils.UUID()
		written := WriteCountry(Country{id, name})
		if written != nil {
			fmt.Printf("[Storage] error inserting country id=%s, name=%s, err=%s\n", id, name, *written)
		}
	}
}
