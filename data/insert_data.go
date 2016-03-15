package data

import (
	"fmt"
	"os"

	"github.com/ivpusic/grpool"
	"github.com/adamdecaf/horizon/data/engines/postgres"
	"github.com/adamdecaf/horizon/data/geo"
)

func InsertData() {
	go insert_raw_data()
	go insert_gzipped_sql()
}

func insert_raw_data() {
	fmt.Println("[data] Starting raw data insert")

	// number of workers, and size of job queue
	pool := grpool.NewPool(10, 10 * 1000 * 1000)

	defer pool.Release()

	if pool == nil {
		fmt.Println("[storage] worker pool is nil")
	} else {
		pool.WaitCount(10)

		if run := os.Getenv("INSERT_RAW_STATES"); run == "yes" {
			if err := geo.InsertRawStates(*pool); err != nil {
				fmt.Printf("[data] Error when inserting raw state data (err=%s)\n", err)
			}
		}

		if run := os.Getenv("INSERT_RAW_CITIES"); run == "yes" {
			if err := geo.InsertRawCitiesFromStates(*pool); err != nil {
				fmt.Printf("[data] Error when inserting raw city data (err=%s)\n", err)
			}
		}

		if run := os.Getenv("INSERT_RAW_COUNTRIES"); run == "yes" {
			if err := geo.InsertCountries(*pool); err != nil {
				fmt.Printf("[data] Error when inserting country data (err=%s)\n", err)
			}
		}

		pool.WaitAll()
	}
}

func insert_gzipped_sql() {
	fmt.Printf("[storage] Inserting gzipped sql files\n")

	if run := os.Getenv("INSERT_HOSTNAMES"); run == "yes" {
                fmt.Printf("[storage] Starting to insert 1m hostnames in .sql file\n")
		rows, err := postgres.ExecuteGzippedSQL("data/raw-data/top-1m-hostnames.sql.gz")
		if err != nil {
			fmt.Printf("[storage] Error when inserting top 1m hostnames from .sql err=%s\n", err)
		}
		fmt.Printf("[storage] Inserted %d hostname rows.\n", rows)
	}
}
