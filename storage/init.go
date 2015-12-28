package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func InitializeStorage() (*sql.DB, error) {
	user := os.Getenv("STORAGE_USER")
	password := os.Getenv("STORAGE_PASSWORD")
	host := os.Getenv("STORAGE_USERNAME")
	dbname := "horizon"

	conn_string := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", conn_string)

	if err != nil {
		LogConnString(conn_string, password)
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func LogConnString(conn_str string, password string) {
	fmt.Println(strings.Replace(conn_str, password, "--password--", 1))
}
