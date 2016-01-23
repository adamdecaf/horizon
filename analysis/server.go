package analysis

import (
	"fmt"
	"net/http"

	"github.com/adamdecaf/horizon/analysis/routes"
)

func StartHttpServer() {
	fmt.Println("[HTTP] Starting http server")

	http.Handle("/", http.FileServer(http.Dir("./analysis/html/")))
	http.HandleFunc("/ping", routes.Ping)

	// search routes
	http.HandleFunc("/cities", routes.SearchCities)
	http.HandleFunc("/humans", routes.SearchHumanNames)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("[HTTP] error when binding and listening: ", err)
	}
}
