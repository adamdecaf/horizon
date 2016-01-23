package routes

import (
	"fmt"
	"net/http"
)

func SearchHumanNames(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "human names")
}
