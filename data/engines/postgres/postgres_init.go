package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	_ "github.com/lib/pq"
)

var db *sql.DB

// todo
// rename method to Init()
// rename file to init.go ?

func InitializePostgres() (*sql.DB, error) {
	// If we have a cache already use it
	if db != nil {
		return db, nil
	}

	user := os.Getenv("STORAGE_USER")
	password := os.Getenv("STORAGE_PASSWORD")
	host := os.Getenv("STORAGE_HOSTNAME")
	port := os.Getenv("STORAGE_PORT")
	dbname := "horizon"

	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	created, err := sql.Open("postgres", conn_string)
	if err != nil {
		LogConnString(conn_string, password)
		log.Fatal(err)
		return nil, err
	}

	db = created

	return db, nil
}

func LogConnString(conn_str string, password string) {
	log.Println(strings.Replace(conn_str, password, "--password--", 1))
}
