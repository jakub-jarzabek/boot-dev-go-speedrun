package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) webhook(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		UserId int `json:"user_id"`
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
    return
	}
	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusOK, "Not a user.upgraded event")
    return
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	_, err = cfg.DB.GetUser(params.Data.UserId)

	if err != nil {
		respondWithError(w, 404, "User not found")
		return
	}

	_, err = cfg.DB.UpgradeUser(params.Data.UserId)

	if err != nil {
		respondWithError(w, 501, "Failed to upgrade user")
		return
	}

	respondWithJSON(w, http.StatusOK, nil)

}
