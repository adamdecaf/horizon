package storage

import (
	"fmt"
	"os"

	"github.com/rubenv/sql-migrate"
)

func InsertData() {
	fmt.Println("[Storage] migrating storage")

	fmt.Println("[Storage] Running .sql insert scripts")
	RunInsertScripts()

	fmt.Println("[Storage] Starting raw data insert")
	InsertRawData()
}

func RunInsertScripts() {
	table_name := "horizon_data_insert"

	db, err := InitializeStorage()
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

func InsertRawData() {
	if run := os.Getenv("INSERT_RAW_STATES"); run == "yes" {
		if err := InsertRawStates(); err != nil {
			fmt.Printf("[Storage/insert] Error when inserting raw state data (err=%s)\n", err)
		}
	}

	if run := os.Getenv("INSERT_RAW_CITIES"); run == "yes" {
		if err := InsertRawCitiesFromStates(); err != nil {
			fmt.Printf("[Storage/insert] Error when inserting raw city data (err=%s)\n", err)
		}
	}

	if run := os.Getenv("INSERT_RAW_COUNTRIES"); run == "yes" {
		if err := InsertCountries(); err != nil {
			fmt.Printf("[Storage/insert] Error when inserting country data (err=%s)\n", err)
		}
	}

	if run := os.Getenv("INSERT_HOSTNAMES"); run == "yes" {
		if err := InsertHostnames(); err != nil {
			fmt.Printf("[Storage/insert] Error when inserting hostnames (err=%s)", err)
		}
	}
}
