package storage

import (
	"fmt"
	"os"

	"github.com/ivpusic/grpool"
	"github.com/rubenv/sql-migrate"
)

func InsertData() {
	go run_insert_scripts()
	go insert_raw_data()
	go insert_gzipped_sql()
}

func run_insert_scripts() {
	fmt.Println("[Storage] Running .sql insert scripts")

	table_name := "horizon_data_insert"

	db, err := InitializePostgres()
	if err != nil {
		panic(err)
	}

	// remove existing rows
	if _, err := db.Query("delete from " + table_name); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Removed rows from " + table_name)

	migrate.SetTable(table_name)
	migrations := &migrate.FileMigrationSource{
		Dir: "storage/migrations/data/",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

func insert_raw_data() {
	fmt.Println("[Storage] Starting raw data insert")

	// number of workers, and size of job queue
	pool := grpool.NewPool(10, 10 * 1000 * 1000)

	defer pool.Release()

	if pool == nil {
		fmt.Println("[storage] worker pool is nil")
	} else {
		pool.WaitCount(10)

		if run := os.Getenv("INSERT_RAW_STATES"); run == "yes" {
			if err := InsertRawStates(*pool); err != nil {
				fmt.Printf("[Storage] Error when inserting raw state data (err=%s)\n", err)
			}
		}

		if run := os.Getenv("INSERT_RAW_CITIES"); run == "yes" {
			if err := InsertRawCitiesFromStates(*pool); err != nil {
				fmt.Printf("[Storage] Error when inserting raw city data (err=%s)\n", err)
			}
		}

		if run := os.Getenv("INSERT_RAW_COUNTRIES"); run == "yes" {
			if err := InsertCountries(*pool); err != nil {
				fmt.Printf("[Storage] Error when inserting country data (err=%s)\n", err)
			}
		}

		pool.WaitAll()
	}
}

func insert_gzipped_sql() {
	fmt.Printf("[storage] Inserting gzipped sql files\n")

	if run := os.Getenv("INSERT_HOSTNAMES"); run == "yes" {
                fmt.Printf("[storage] Starting to insert 1m hostnames in .sql file\n")
		rows, err := ExecuteGzippedSQL("storage/raw-data/top-1m-hostnames.sql.gz")
		if err != nil {
			fmt.Printf("[storage] Error when inserting top 1m hostnames from .sql err=%s\n", err)
		}
		fmt.Printf("[storage] Inserted %d hostname rows.\n", rows)
	}
}
