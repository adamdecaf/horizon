package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamdecaf/horizon/storage"
)

type CityState struct {
	City storage.City `json:"city"`
	State storage.State `json:"state"`
}

type CityResult struct {
	CityStates []CityState `json:"cities"`
}

func SearchCities(w http.ResponseWriter, r *http.Request) {
	var results []CityState

	query := r.URL.Query().Get("q")
	if query != "" {
		cities, err := storage.SearchCitiesByName(query)
		if err != nil {
			fmt.Printf("[Analysis] error getting cities param='%s' err='%s'\n", query, err)
			w.WriteHeader(503)
			return
		}

		for i := range cities {
			state, err := find_state(cities[i].StateId)
			if err != nil {
				fmt.Printf("[Analysis] error finding state from cache param='%s' err='%s'\n", query, err)
				w.WriteHeader(503)
				return
			}

			results = append(results, CityState{cities[i], state})
		}

		res, err := json.Marshal(CityResult{results})
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

func find_state(state_id string) (storage.State, error) {
	state, err := storage.FindStateById(state_id)
	if err != nil {
		return storage.State{}, err
	}

	if state != nil {
		return *state, nil
	}

	return storage.State{}, fmt.Errorf("unable to really find state... (state_id=%s)", state_id)
}
