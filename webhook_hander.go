package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) webhook(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		UserId string `json:"user_id"`
	}

	type parameters struct {
		Event string `json:"event"`
		Data  Data   `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	}
	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusOK, "Not a user.upgraded event")
	}

	userId, err := strconv.Atoi(params.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	_, err = cfg.DB.GetUser(userId)

	if err != nil {
		respondWithError(w, 404, "User not found")
		return
	}

	respondWithJSON(w, http.StatusOK, nil)

}
