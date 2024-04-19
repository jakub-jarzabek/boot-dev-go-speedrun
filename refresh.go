package main

import (
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	log.Print("handling handlerRefresh")
	authHeader := r.Header.Get("Authorization")
	token := strings.Split(authHeader, " ")[1]

	log.Print(token)
	id, err := validateRefreshToken(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}
	revoked, err := cfg.DB.IsTokenRevoked(token)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}
	if revoked {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	newToken, err := generateAccessToken(id, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}
	type Token struct {
		Token string `json:"token"`
	}

	log.Print(newToken)

	respondWithJSON(w, 200, Token{Token: newToken})
}
