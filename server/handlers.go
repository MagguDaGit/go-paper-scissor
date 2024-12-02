package server

import (
	"encoding/json"
	"net/http"
	"rubrumcreation.com/go-paper-scissor/services"
)

func randomHandler(w http.ResponseWriter, r *http.Request, title string) {
	numOfGames := 10
	summary := services.PlayRandomGames(numOfGames)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
