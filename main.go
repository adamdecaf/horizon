package main

import (
	"fmt"

	"github.com/adamdecaf/horizon/storage"
)

func main() {
	fmt.Println("Starting horizon")

	storage.MigrateStorage()
}
