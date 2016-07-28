package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	_ "github.com/lib/pq"
	"github.com/adamdecaf/horizon/utils"
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

	config := utils.NewConfig()

	user := config.Get("STORAGE_USER")
	password := config.Get("STORAGE_PASSWORD")
	host := config.Get("STORAGE_HOSTNAME")
	port := config.Get("STORAGE_PORT")
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
