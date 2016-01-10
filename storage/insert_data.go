package storage

import (
	"fmt"

	"github.com/rubenv/sql-migrate"
)

func InsertData() {
	fmt.Println("[Storage] migrating storage")

	fmt.Println("[Storage] Running .sql insert scripts")
	RunInsertScripts()

	fmt.Println("[Storage] Inserting raw data")
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
	states, err := InsertRawStates()
	if err != nil {
		fmt.Printf("[Storage/insert] Error when inserting raw state data (err=%s)\n", err)
	} else {
		fmt.Printf("[Storage/insert] Inserted %d states\n", len(states))
	}

	cities_count, err := InsertRawCitiesFromStates(states)
	if err != nil {
		fmt.Printf("[Storage/insert] Error when inserting raw state data (err=%s)\n", err)
	} else {
		fmt.Printf("[Storage/insret] Inserted %d cities\n", cities_count)
	}
}
