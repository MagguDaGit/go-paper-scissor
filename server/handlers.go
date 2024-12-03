package server

import (
	"encoding/json"
	"net/http"
	"net/url"

	"rubrumcreation.com/go-paper-scissor/models"
	"rubrumcreation.com/go-paper-scissor/services"
	"rubrumcreation.com/go-paper-scissor/util"
)

func randomSimulateHandler(w http.ResponseWriter, r *http.Request, query url.Values) {
	// get key,and validate presence of expected query key/value
	paramValue, err := util.GetExpectedParam("amount", query)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	// Convert and validate the expected value
	numOfGames, err := util.ConvertInt(paramValue)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Call service for application logic
	summary := services.PlayRandomGames(numOfGames)
	w.Header().Set("Content-Type", "application/json")

	// convert to json
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func playRandomHandler(w http.ResponseWriter, r *http.Request, query url.Values) {
	expectedMoveKey := "move"
	expectedNameKey := "playerName"

	paramValue, err := util.GetExpectedParam(expectedMoveKey, query)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	paramNameValue, err := util.GetExpectedParam(expectedNameKey, query)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	move, err := util.ConvertMove(paramValue)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	p1 := models.GenericPlayer{}
	p1.Name = models.PlayerName(paramNameValue)
	p1.ChosenMove = move

	p2 := models.RandomPlayer{Name: models.PLAYER_2}
	summary := services.Play(p1, p2)

	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
