package analysis

import (
	"net/http"
	"log"
	"github.com/adamdecaf/horizon/analysis/routes"
)

func StartHttpServer() {
	log.Println("[HTTP] Starting http server")

	http.Handle("/", http.FileServer(http.Dir("./analysis/html/")))
	http.HandleFunc("/ping", routes.Ping)

	// search routes
	http.HandleFunc("/cities", routes.SearchCities)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("[HTTP] error when binding and listening: ", err)
	}
}
