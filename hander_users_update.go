package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	authHeader := r.Header.Get("Authorization")
	token := strings.Split(authHeader, " ")[1]

	id, err := validateAccessToken(token, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")

	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.UpdateUser(id, params.Email, params.Password)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, 200, User{
		ID:    user.ID,
		Email: user.Email,
	})
}
