package main

import (
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.Split(authHeader, " ")[1]
	_, err := validateRefreshToken(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	success, err := cfg.DB.AddRevokedToken(token)
	if err != nil {
		respondWithError(w, 500, "Couldn't revoke token")
		return
	}
	respondWithJSON(w, 200, success)

}
