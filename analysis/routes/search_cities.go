package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamdecaf/horizon/storage"
)

type CityResult struct {
	Cities []storage.City `json:"cities"`
}

func SearchCities(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query != "" {
		cities, err := storage.SearchCitiesByName(query)
		if err != nil {
			fmt.Printf("[Analysis] error getting cities param='%s' err='%s'\n", query, err)
			w.WriteHeader(503)
		}

		res, err := json.Marshal(CityResult{cities})
		if err != nil {
			fmt.Printf("[Analysis] error marshalling cities param='%s' err='%s'\n", query, err)
			w.WriteHeader(503)
		} else {
			w.Write(res)
		}

	} else {
		w.WriteHeader(204)
	}
}
