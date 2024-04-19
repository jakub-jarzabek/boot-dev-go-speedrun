package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {

	token, err := getAuthToken(r)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")

	}

	id, err := validateAccessToken(token, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")

	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	if dbChirp.AuthorID != id {
		respondWithError(w, http.StatusForbidden, "You can't delete this chirp")
		return
	}

	cfg.DB.DeleteChirp(chirpID)

	respondWithJSON(w, http.StatusOK, "OK")
}
