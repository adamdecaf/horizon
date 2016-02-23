package storage

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

func MigrateStorage() {
	fmt.Println("migrating storage")

	migrations := &migrate.FileMigrationSource{
		Dir: "storage/migrations/",
	}

	db, err := InitializePostgres()
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
